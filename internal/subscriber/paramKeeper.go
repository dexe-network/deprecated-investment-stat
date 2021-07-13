package subscriber

import (
	"dex-trades-parser/internal/models"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"time"
)

func (s *Subscriber) paramKeeperTransactionProcessing(tx types.Transaction, blockNumber int64, blockTime uint64) {
	_, data, method, err := s.parser.ParamKeeperTransactionInfo(tx)

	if err != nil {
		s.log.Debug("Cent Parse Tx: "+tx.Hash().String(), zap.Error(err))
		return
	}

	switch method.Name {

	case "whitelistToken":

		err = s.st.Repo.GlobalTokenWhitelist.Save(&models.GlobalTokenWhitelist{
			Address:     data["_token"].(common.Address).String(),
			Date:        time.Unix(int64(blockTime), 0),
			BlockNumber: blockNumber,
			Tx:          tx.Hash().String(),
		})
		if err != nil {
			s.log.Error("Can't save whitelistToken to DB", zap.Error(err))
		}
	case "delistToken":

		err = s.st.Repo.GlobalTokenWhitelist.Delete(data["_token"].(common.Address).String())
		if err != nil {
			s.log.Error("Can't Delete delistToken from DB", zap.Error(err))
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
