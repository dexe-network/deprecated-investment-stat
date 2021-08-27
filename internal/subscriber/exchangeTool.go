package subscriber

import (
	"dex-trades-parser/internal/contracts/traderPoolUpgradeable"
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (s *Subscriber) exchangeToolTransactionProcessing(tx types.Transaction, blockNumber int64, blockTime uint64) {
	parsedTransaction, err := s.parser.ParseExchangeToolTransaction(tx)
	if err != nil {
		s.log.Debug("Cent Parse Tx: "+tx.Hash().String(), zap.Error(err))
		return
	}
	parsedPaths := &pgtype.TextArray{}
	parsedPaths.Set(helpers.AddressArrToStringArr(parsedTransaction.Path))

	// Get Trader Wallet Address
	var foundPool models.Pool
	if err := s.st.DB.First(&foundPool, "LOWER(\"poolAdr\") = LOWER(?)", parsedTransaction.TraderPool.String()).
		Error; err != nil {
		s.log.Debug("Find Pool error "+tx.Hash().String(), zap.Error(err))
		return
	}

	var tradeType string
	if parsedTransaction.Path[0].String() == foundPool.BasicTokenAdr {
		tradeType = "buy"
	} else if parsedTransaction.Path[len(parsedTransaction.Path)-1].String() == foundPool.BasicTokenAdr {
		tradeType = "sell"
	}

	err = s.st.Repo.Trade.Save(
		&models.Trade{
			TraderPool:   parsedTransaction.TraderPool.String(),
			AmountIn:     parsedTransaction.AmountIn.String(),
			AmountOutMin: parsedTransaction.AmountOutMin.String(),
			Path:         *parsedPaths,
			Deadline:     parsedTransaction.Deadline.String(),
			Date:         time.Unix(int64(blockTime), 0),
			Type:         tradeType,
			BlockNumber:  parsedTransaction.BlockNumber,
			Tx:           parsedTransaction.Tx,
		},
	)
	if err != nil {
		s.log.Error("Can't save trade to DB", zap.Error(err))
	}

	///////// Store Pool Indicators for every operation
	instance, err := traderPoolUpgradeable.NewToken(parsedTransaction.TraderPool, s.client)
	if err != nil {
		s.log.Debug("Create instance of token error "+tx.Hash().String(), zap.Error(err))
		return
	}

	totalCap, totalSupply, err := instance.GetTotalValueLocked(&bind.CallOpts{BlockNumber: big.NewInt(parsedTransaction.BlockNumber)})
	if err != nil {
		s.log.Debug("GetTotalValueLocked request error "+tx.Hash().String(), zap.Error(err))
		return
	}
	poolTokenPrice := decimal.NewFromInt(0)
	if totalCap.Cmp(big.NewInt(0)) > 0 && totalSupply.Cmp(big.NewInt(0)) > 0 {
		totalCapDecimal := helpers.ToDecimal(totalCap, int(foundPool.BasicTokenDecimals))
		totalSupplyDecimal := helpers.ToDecimal(totalSupply, int(foundPool.Decimals))
		poolTokenPrice = totalCapDecimal.Div(totalSupplyDecimal)
	}

	traderAmount, _, _, err := instance.GetUserData(
		&bind.CallOpts{BlockNumber: big.NewInt(parsedTransaction.BlockNumber)},
		common.HexToAddress(foundPool.CreatorAdr),
	)
	if err != nil {
		s.log.Debug("GetUserData request error "+tx.Hash().String(), zap.Error(err))
		return
	}

	err = s.st.Repo.PoolIndicators.Save(
		&models.PoolIndicators{
			PoolAdr:                    parsedTransaction.TraderPool.String(),
			TotalCap:                   totalCap.String(),
			TotalSupply:                totalSupply.String(),
			TraderBasicTokensDeposited: traderAmount.String(),
			PoolTokenPrice:             poolTokenPrice.String(),
			Date:                       time.Unix(int64(blockTime), 0),
			BlockNumber:                parsedTransaction.BlockNumber,
			Tx:                         parsedTransaction.Tx,
		},
	)
	if err != nil {
		s.log.Error("Can't save PoolIndicators to DB", zap.Error(err))
	}
	/////////
}
