package subscriber

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgtype"
	"go.uber.org/zap"
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
		AmountIn:     pgtype.Numeric{Int: parsedTransaction.AmountIn, Status: pgtype.Present},
		AmountOutMin: pgtype.Numeric{Int: parsedTransaction.AmountOutMin, Status: pgtype.Present},
		Path:         *parsedPaths,
		Deadline:     pgtype.Numeric{Int: parsedTransaction.Deadline, Status: pgtype.Present},
		Date:         time.Unix(int64(blockTime), 0),
		BlockNumber:  parsedTransaction.BlockNumber,
		Tx:           parsedTransaction.Tx,
	})
	if err != nil {
		s.log.Error("Can't save trade to DB", zap.Error(err))
	}
}
