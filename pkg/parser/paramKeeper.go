package parser

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

func (p *Parser) ParamKeeperTransactionInfo(t types.Transaction) (receipt *types.Receipt, data map[string]interface{}, method *abi.Method, err error) {

	receipt, err = p.client.TransactionReceipt(context.Background(), t.Hash())
	if err != nil {
		return
	} else {
		if receipt.Status != 1 {
			err = errors.New("transaction status fail : " + t.Hash().String())
			return
		}
	}

	if len(t.Data()) < 5 {
		err = errors.New("transaction data to small")
		p.log.Debug("Tx Data len < 5", zap.Error(err), zap.String("Tx", t.Hash().String()))
		return
	}

	data = make(map[string]interface{})
	method, err = p.Abis.ParamKeeper.MethodById(t.Data()[:4])
	if err != nil {
		p.log.Debug("Tx Data len < 5", zap.Error(err), zap.String("Tx", t.Hash().String()))
		return
	}
	err = method.Inputs.UnpackIntoMap(data, t.Data()[4:])
	if err != nil {
		p.log.Debug("UnpackIntoMap", zap.Error(err), zap.String("Tx", t.Hash().String()))
		return
	}

	return
}
