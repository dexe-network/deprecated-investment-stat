package parser

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"math/big"
)

type ParsedExchangeToolTx struct {
	TraderPool   common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	Deadline     *big.Int
	BlockNumber  int64
	Tx           string
}

func (p *Parser) ParseExchangeToolTransaction(t types.Transaction) (pTx ParsedExchangeToolTx, err error) {

	receipt, err := p.client.TransactionReceipt(context.Background(), t.Hash())
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

	data := make(map[string]interface{})
	method, err := p.Abis.ExchangeTool.MethodById(t.Data()[:4])
	if err != nil {
		p.log.Debug("Tx Data len < 5", zap.Error(err), zap.String("Tx", t.Hash().String()))
		return
	}
	err = method.Inputs.UnpackIntoMap(data, t.Data()[4:])
	if err != nil {
		p.log.Debug("UnpackIntoMap", zap.Error(err), zap.String("Tx", t.Hash().String()))
		return
	}

	switch method.Name {

	case "swapExactTokensForTokens":
		pTx.TraderPool = data["traderPool"].(common.Address)
		pTx.AmountIn, _ = data["amountIn"].(*big.Int)
		pTx.AmountOutMin, _ = data["amountOutMin"].(*big.Int)
		pTx.Path = data["path"].([]common.Address)
		pTx.Deadline, _ = data["deadline"].(*big.Int)
		pTx.Tx = t.Hash().String()
		pTx.BlockNumber = receipt.BlockNumber.Int64()

	default:
		fmt.Println("Unknown metod : " + method.Name)
		err = errors.New("Unknown metod : " + method.Name)
	}

	return
}
