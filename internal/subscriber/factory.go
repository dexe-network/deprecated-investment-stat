package subscriber

import (
	"dex-trades-parser/internal/models"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgtype"
	"go.uber.org/zap"
	"time"
)

func (s *Subscriber) factoryTransactionProcessing(tx types.Transaction, blockNumber int64, blockTime uint64) {
	parsedTransaction, err := s.parser.ParseFactoryTransaction(tx)
	if err != nil {
		s.log.Debug("Cent Parse Tx: "+tx.Hash().String(), zap.Error(err))
		return
	}

	err = s.st.Repo.Pool.Save(&models.Pool{
		CreatorAdr:            parsedTransaction.CreatorAdr.String(),
		BasicTokenAdr:         parsedTransaction.BasicTokenAdr.String(),
		TotalSupply:           pgtype.Numeric{Int: parsedTransaction.TotalSupply, Status: pgtype.Present},
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
		Date:                  time.Unix(int64(blockTime), 0),
		BlockNumber:           parsedTransaction.BlockNumber,
		Tx:                    parsedTransaction.Tx,
	})
	if err != nil {
		s.log.Error("Can't save pool to DB", zap.Error(err))
	}
}
