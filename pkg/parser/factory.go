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

type ParsedFactoryTx struct {
	CreatorAdr            common.Address
	BasicTokenAdr         common.Address
	TotalSupply           *big.Int
	TraderCommissionNum   uint16
	TraderCommissionDen   uint16
	InvestorCommissionNum uint16
	InvestorCommissionDen uint16
	DexeCommissionNum     uint16
	DexeCommissionDen     uint16
	IsActualOn            bool
	InvestorRestricted    bool
	Name                  string
	Symbol                string
	PoolAdr               string
	BlockNumber           int64
	Tx                    string
}

func (p *Parser) ParseFactoryTransaction(t types.Transaction) (pTx ParsedFactoryTx, err error) {

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
	method, err := p.Abis.TraderPoolFactory.MethodById(t.Data()[:4])
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

	case "createTraderContract":
		commisons := data["_comm"].([6]uint16)
		pTx.CreatorAdr = data["_traderWallet"].(common.Address)
		pTx.BasicTokenAdr = data["_basicToken"].(common.Address)
		pTx.TotalSupply = data["_totalSupply"].(*big.Int)
		pTx.TraderCommissionNum = commisons[0]
		pTx.TraderCommissionDen = commisons[1]
		pTx.InvestorCommissionNum = commisons[2]
		pTx.InvestorCommissionDen = commisons[3]
		pTx.DexeCommissionNum = commisons[4]
		pTx.DexeCommissionDen = commisons[5]
		pTx.IsActualOn = data["_actual"].(bool)
		pTx.InvestorRestricted = data["_investorRestricted"].(bool)
		pTx.Name = data["_name"].(string)
		pTx.Symbol = data["_symbol"].(string)
		pTx.Tx = t.Hash().String()
		pTx.BlockNumber = receipt.BlockNumber.Int64()
		pTx.PoolAdr = common.BytesToAddress(receipt.Logs[len(receipt.Logs)-1].Data).String()

	default:
		fmt.Println("Unknown metod : " + method.Name)
		err = errors.New("Unknown metod : " + method.Name)
	}

	return
}
