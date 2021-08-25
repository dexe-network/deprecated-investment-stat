package subscriber

import (
	"dex-trades-parser/internal/contracts/erc20"
	"dex-trades-parser/internal/contracts/traderPoolUpgradeable"
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (s *Subscriber) factoryTransactionProcessing(tx types.Transaction, blockNumber int64, blockTime uint64) {
	parsedTransaction, err := s.parser.ParseFactoryTransaction(tx)
	if err != nil {
		s.log.Debug("Cent Parse Tx: "+tx.Hash().String(), zap.Error(err))
		return
	}

	// Get Basic Token Info
	instanceErc20, err := erc20.NewErc20(parsedTransaction.BasicTokenAdr, s.client)
	if err != nil {
		s.log.Debug("Create instance of token error "+tx.Hash().String(), zap.Error(err))
		return
	}

	basicTokenDecimals, err := instanceErc20.Decimals(&bind.CallOpts{})
	if err != nil {
		s.log.Debug("Decimals request error "+tx.Hash().String(), zap.Error(err))
		return
	}

	basicTokenSymbol, err := instanceErc20.Symbol(&bind.CallOpts{})
	if err != nil {
		s.log.Debug("Decimals request error "+tx.Hash().String(), zap.Error(err))
		return
	}
	//

	err = s.st.Repo.Pool.Save(&models.Pool{
		CreatorAdr:            parsedTransaction.CreatorAdr.String(),
		BasicTokenAdr:         parsedTransaction.BasicTokenAdr.String(),
		BasicTokenDecimals:    basicTokenDecimals,
		BasicTokenSymbol:      basicTokenSymbol,
		TotalSupply:           parsedTransaction.TotalSupply.String(),
		TraderCommissionNum:   parsedTransaction.TraderCommissionNum,
		TraderCommissionDen:   parsedTransaction.TraderCommissionDen,
		InvestorCommissionNum: parsedTransaction.InvestorCommissionNum,
		InvestorCommissionDen: parsedTransaction.InvestorCommissionDen,
		DexeCommissionNum:     parsedTransaction.DexeCommissionNum,
		DexeCommissionDen:     parsedTransaction.DexeCommissionDen,
		IsActualOn:            parsedTransaction.IsActualOn,
		InvestorRestricted:    parsedTransaction.InvestorRestricted,
		Name:                  parsedTransaction.Name,
		Symbol:                parsedTransaction.Symbol,
		PoolAdr:               parsedTransaction.PoolAdr,
		Decimals:              uint8(18), // Fixed Decimal, need improve after upgrade smart contracts
		Date:                  time.Unix(int64(blockTime), 0),
		BlockNumber:           parsedTransaction.BlockNumber,
		Tx:                    parsedTransaction.Tx,
	})
	if err != nil {
		s.log.Error("Can't save pool to DB", zap.Error(err))
	}

	///////// Store Pool Indicators for every operation
	instance, err := traderPoolUpgradeable.NewToken(common.HexToAddress(parsedTransaction.PoolAdr), s.client)
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
		totalCapDecimal := helpers.ToDecimal(totalCap, int(basicTokenDecimals))
		totalSupplyDecimal := helpers.ToDecimal(totalSupply, int(uint8(18)))
		poolTokenPrice = totalCapDecimal.Div(totalSupplyDecimal)
	}

	traderAmount, _, _, err := instance.GetUserData(
		&bind.CallOpts{BlockNumber: big.NewInt(parsedTransaction.BlockNumber)}, parsedTransaction.CreatorAdr)
	if err != nil {
		s.log.Debug("GetUserData request error "+tx.Hash().String(), zap.Error(err))
		return
	}

	err = s.st.Repo.PoolIndicators.Save(&models.PoolIndicators{
		PoolAdr:                    parsedTransaction.PoolAdr,
		TotalCap:                   totalCap.String(),
		TotalSupply:                totalSupply.String(),
		TraderBasicTokensDeposited: traderAmount.String(),
		PoolTokenPrice:             poolTokenPrice.String(),
		Date:                       time.Unix(int64(blockTime), 0),
		BlockNumber:                parsedTransaction.BlockNumber,
		Tx:                         parsedTransaction.Tx,
	})
	if err != nil {
		s.log.Error("Can't save PoolIndicators to DB", zap.Error(err))
	}
	/////////
}
