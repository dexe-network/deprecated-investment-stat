package subscriber

import (
	"dex-trades-parser/internal/contracts/traderPoolUpgradeable"
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgtype"
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

	err = s.st.Repo.Trade.Save(&models.Trade{
		TraderPool:   parsedTransaction.TraderPool.String(),
		AmountIn:     parsedTransaction.AmountIn.String(),
		AmountOutMin: parsedTransaction.AmountOutMin.String(),
		Path:         *parsedPaths,
		Deadline:     parsedTransaction.Deadline.String(),
		Date:         time.Unix(int64(blockTime), 0),
		BlockNumber:  parsedTransaction.BlockNumber,
		Tx:           parsedTransaction.Tx,
	})
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

	err = s.st.Repo.PoolIndicators.Save(&models.PoolIndicators{
		PoolAdr:     parsedTransaction.TraderPool.String(),
		TotalCap:    totalCap.String(),
		TotalSupply: totalSupply.String(),
		Date:        time.Unix(int64(blockTime), 0),
		BlockNumber: parsedTransaction.BlockNumber,
		Tx:          parsedTransaction.Tx,
	})
	if err != nil {
		s.log.Error("Can't save PoolIndicators to DB", zap.Error(err))
	}
	/////////
}
