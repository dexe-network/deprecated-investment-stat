package subscriber

import (
	"dex-trades-parser/internal/models"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (s *Subscriber) traderPoolTransactionProcessing(tx types.Transaction, blockNumber int64, blockTime uint64) {
	_, data, method, err := s.parser.BaseTraderPoolTransactionInfo(tx)

	if err != nil {
		s.log.Debug("Cent Parse Tx: "+tx.Hash().String(), zap.Error(err))
		return
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
	if err != nil {
		s.log.Debug("Cent Parse AsMessage: "+tx.Hash().String(), zap.Error(err))
		return
	}

	switch method.Name {

	case "deposit":
		err = s.st.Repo.PoolTransfer.Save(&models.PoolTransfer{
			Wallet:      msg.From().String(),
			PoolAdr:     tx.To().String(),
			Amount:      data["amount"].(*big.Int).String(),
			Type:        "deposit",
			Date:        time.Unix(int64(blockTime), 0),
			BlockNumber: blockNumber,
			Tx:          tx.Hash().String(),
		})
		if err != nil {
			s.log.Error("Can't save trade to DB", zap.Error(err))
		}
	case "withdraw":

		err = s.st.Repo.PoolTransfer.Save(&models.PoolTransfer{
			Wallet:      msg.From().String(),
			PoolAdr:     tx.To().String(),
			Amount:      data["amount"].(*big.Int).String(),
			Type:        "withdraw",
			Date:        time.Unix(int64(blockTime), 0),
			BlockNumber: blockNumber,
			Tx:          tx.Hash().String(),
		})
		if err != nil {
			s.log.Error("Can't save trade to DB", zap.Error(err))
		}

	default:
		fmt.Println("Unknown metod : " + method.Name)
		err = errors.New("Unknown metod : " + method.Name)
	}

	if err != nil {
		s.log.Debug("Cent Parse Tx: "+tx.Hash().String(), zap.Error(err))
		return
	}
}
