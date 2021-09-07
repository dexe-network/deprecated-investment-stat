package parser

import (
	"context"
	"dex-trades-parser/internal/models"
	"dex-trades-parser/internal/storage"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
	"math/big"
	"time"
)

type ParsedExchangeToolTx struct {
	TraderPool common.Address
	AmountIn *big.Int
	AmountOutMin *big.Int
	FromAmt *big.Int
	ToAmt *big.Int
	Path []common.Address
	Deadline *big.Int
	BlockNumber int64
	Tx string
}

type ExchangeLogData struct {
	FromAsset common.Address
	ToAsset   common.Address
	FromAmt   *big.Int
	ToAmt     *big.Int
}

func (p *Parser) CalculateTrades(data *ParsedExchangeToolTx, tradeType string, blockTime uint64, st *storage.Storage) (err error) {
	var tokenAddress common.Address
	if tradeType == "buy" {
		tokenAddress = data.Path[len(data.Path)-1]
	} else if tradeType == "sell" {
		tokenAddress = data.Path[0]
	}

	var tradeItems []models.TradeItem

	if err := st.DB.Order("date desc").Find(
		&tradeItems,
		"LOWER(\"poolAddress\") = LOWER(?) AND LOWER(\"tokenAddress\") = LOWER(?) AND balance > 0",
		data.TraderPool.String(),
		tokenAddress.String(),
	).Error;
		err != nil {
		p.log.Debug("Error CalculateTrades Find tradeItem", zap.Error(err))
		return err
	}

	if len(tradeItems) <= 0 {

		newTradeItem := models.TradeItem{
			PoolAddress:  data.TraderPool.String(),
			TokenAddress: tokenAddress.String(),
			Balance:      data.ToAmt.String(),
			TradeStatus:  "open",
			TradeEvents:  nil,
			OpenDate:     time.Unix(int64(blockTime), 0),
			CloseDate:    time.Time{},
			Date:         time.Unix(int64(blockTime), 0),
			BlockNumber:  data.BlockNumber,
		}

		newTradeItem.TradeEvents = append(newTradeItem.TradeEvents, models.TradeEvent{
			TradeType:   tradeType,
			Amount:      data.ToAmt.String(),
			Date:        time.Unix(int64(blockTime), 0),
			BlockNumber: data.BlockNumber,
			Tx:          data.Tx,
		})

		if err := st.DB.Create(&newTradeItem).Error; err != nil {
			p.log.Debug("Error Create newTradeItem", zap.Error(err))
			return err
		}
	} else {
		openTrade := tradeItems[0]

		balance := new(big.Int)
		balance.SetString(openTrade.Balance, 10)

		var balanceAfterOperation *big.Int
		var amount string
		if tradeType == "buy" {
			balanceAfterOperation = balance.Add(balance, data.ToAmt)
			amount = data.ToAmt.String()
		} else if tradeType == "sell" {
			balanceAfterOperation = balance.Sub(balance, data.FromAmt)
			amount = data.FromAmt.String()
		}

		if balanceAfterOperation.Cmp(big.NewInt(0)) <= 0 {
			openTrade.TradeStatus = "close"
		}

		openTrade.Balance = balanceAfterOperation.String()

		openTrade.TradeEvents = append(openTrade.TradeEvents, models.TradeEvent{
			TradeType:   tradeType,
			Amount:      amount,
			Date:        time.Unix(int64(blockTime), 0),
			BlockNumber: data.BlockNumber,
			Tx:          data.Tx,
		})
		st.DB.Save(&openTrade)
	}

	return err
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

	// Parse Logs
	var logExchangedData ExchangeLogData
	logExchangedSig := []byte("Exchanged(address,address,uint256,uint256)")
	logExchangedSigHash := crypto.Keccak256Hash(logExchangedSig)
	for _, vLog := range receipt.Logs {
		switch vLog.Topics[0].Hex() {
		case logExchangedSigHash.Hex():
			logExchangedRawData, err := p.Abis.TraderPool.Unpack("Exchanged", vLog.Data)
			if err != nil {
				break
			}

			logExchangedData.FromAsset = logExchangedRawData[0].(common.Address)
			logExchangedData.ToAsset = logExchangedRawData[1].(common.Address)
			logExchangedData.FromAmt = logExchangedRawData[2].(*big.Int)
			logExchangedData.ToAmt = logExchangedRawData[3].(*big.Int)
		}
	}

	if err != nil {
		p.log.Debug("Exchanged", zap.Error(err), zap.String("Tx", t.Hash().String()))
		return
	}

	pTx.FromAmt = logExchangedData.FromAmt
	pTx.ToAmt = logExchangedData.ToAmt

	return
}
