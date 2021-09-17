package parser

import (
	"context"
	"dex-trades-parser/internal/contracts/erc20"
	"dex-trades-parser/internal/contracts/uniPair"
	"dex-trades-parser/internal/models"
	"dex-trades-parser/internal/storage"
	"dex-trades-parser/pkg/helpers"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"strings"
	"time"
)

type ParsedExchangeToolTx struct {
	TraderPool   common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	FromAmt      *big.Int
	ToAmt        *big.Int
	Path         []common.Address
	Deadline     *big.Int
	BlockNumber  int64
	Tx           string
}

type ExchangeLogData struct {
	FromAsset common.Address
	ToAsset   common.Address
	FromAmt   *big.Int
	ToAmt     *big.Int
}

func (p *Parser) CalculateTrades(
	data *ParsedExchangeToolTx,
	tradeType string,
	blockTime uint64,
	st *storage.Storage,
) (err error) {
	fromTokenAddress := data.Path[0]
	toTokenAddress := data.Path[len(data.Path)-1]

	var baseTokenAddress common.Address
	var tradeTokenAddress common.Address

	if tradeType == "buy" {
		baseTokenAddress = fromTokenAddress
		tradeTokenAddress = toTokenAddress
	} else if tradeType == "sell" {
		baseTokenAddress = toTokenAddress
		tradeTokenAddress = fromTokenAddress
	}

	// Trade Token Info
	fromTokenErc20, err := erc20.NewErc20(fromTokenAddress, p.client)
	if err != nil {
		p.log.Debug("Create instance of token error", zap.Error(err))
		return
	}

	fromTokenDecimals, err := fromTokenErc20.Decimals(&bind.CallOpts{})
	if err != nil {
		p.log.Debug("Decimals request error ", zap.Error(err))
		return
	}

	toTokenErc20, err := erc20.NewErc20(toTokenAddress, p.client)
	if err != nil {
		p.log.Debug("Create instance of token error", zap.Error(err))
		return
	}

	toTokenDecimals, err := toTokenErc20.Decimals(&bind.CallOpts{})
	if err != nil {
		p.log.Debug("Decimals request error ", zap.Error(err))
		return
	}
	//

	var tradeItems []models.TradeItem

	if err := st.DB.Order("date desc").Find(
		&tradeItems,
		"LOWER(\"poolAddress\") = LOWER(?) AND LOWER(\"tradeTokenAddress\") = LOWER(?) AND balance > 0",
		data.TraderPool.String(),
		tradeTokenAddress.String(),
	).Error;
		err != nil {
		p.log.Debug("Error CalculateTrades Find tradeItem", zap.Error(err))
		return err
	}

	//sendValue, receiveValue := decimalNumberShifter(data.FromAmt, data.ToAmt, fromTokenDecimals, toTokenDecimals)
	//price := decimal.NewFromBigInt(sendValue, 0).Div(decimal.NewFromBigInt(receiveValue, 0))

	sendValue := helpers.ToDecimal(data.FromAmt, int(fromTokenDecimals))
	receiveValue := helpers.ToDecimal(data.ToAmt, int(toTokenDecimals))

	if len(tradeItems) <= 0 {

		newTradeItem := models.TradeItem{
			PoolAddress:       data.TraderPool.String(),
			BaseTokenAddress:  baseTokenAddress.String(),
			TradeTokenAddress: tradeTokenAddress.String(),
			Balance:           data.ToAmt.String(),
			TradeStatus:       "open",
			TradeEvents:       nil,
			OpenDate:          time.Unix(int64(blockTime), 0),
			CloseDate:         time.Time{},
			Date:              time.Unix(int64(blockTime), 0),
			BlockNumber:       data.BlockNumber,
		}

		price := sendValue.Div(receiveValue)
		//price := decimal.NewFromBigInt(sendValue, 0).Div(decimal.NewFromBigInt(receiveValue, 0))

		newTradeItem.TradeEvents = append(
			newTradeItem.TradeEvents, models.TradeEvent{
				TradeType:            tradeType,
				FromAmount:           data.FromAmt.String(),
				ToAmount:             data.ToAmt.String(),
				Price:                price.String(),
				SpentTokenAddress:    toTokenAddress.String(),
				ReceivedTokenAddress: fromTokenAddress.String(),
				Date:                 time.Unix(int64(blockTime), 0),
				BlockNumber:          data.BlockNumber,
				Tx:                   data.Tx,
			},
		)

		if err := st.DB.Create(&newTradeItem).Error; err != nil {
			p.log.Debug("Error Create newTradeItem", zap.Error(err))
			return err
		}
	} else {
		openTrade := tradeItems[0]

		balance := new(big.Int)
		balance.SetString(openTrade.Balance, 10)

		var balanceAfterOperation *big.Int
		var price decimal.Decimal

		if tradeType == "buy" {
			balanceAfterOperation = balance.Add(balance, data.ToAmt)
			price = sendValue.Div(receiveValue)
			//price = decimal.NewFromBigInt(sendValue, 0).Div(decimal.NewFromBigInt(receiveValue, 0))
		} else if tradeType == "sell" {
			balanceAfterOperation = balance.Sub(balance, data.FromAmt)
			price = receiveValue.Div(sendValue)
			//price = decimal.NewFromBigInt(receiveValue, 0).Div(decimal.NewFromBigInt(sendValue, 0))
		}

		if balanceAfterOperation.Cmp(big.NewInt(0)) <= 0 {
			openTrade.TradeStatus = "close"
		}

		openTrade.Balance = balanceAfterOperation.String()

		openTrade.TradeEvents = append(
			openTrade.TradeEvents, models.TradeEvent{
				TradeType:            tradeType,
				FromAmount:           data.FromAmt.String(),
				ToAmount:             data.ToAmt.String(),
				Price:                price.String(),
				SpentTokenAddress:    toTokenAddress.String(),
				ReceivedTokenAddress: fromTokenAddress.String(),
				Date:                 time.Unix(int64(blockTime), 0),
				BlockNumber:          data.BlockNumber,
				Tx:                   data.Tx,
			},
		)
		st.DB.Save(&openTrade)
	}

	return err
}

func decimalNumberShifter(number0 *big.Int, number1 *big.Int, decimal0 uint8, decimal1 uint8) (
	resultNum0 *big.Int,
	resultNum1 *big.Int,
) {
	if decimal0 < decimal1 {
		shiftDecimal := big.NewInt(int64(decimal1 - decimal0))
		multi := big.NewInt(10)
		modifiedNumber0 := big.NewInt(0).Mul(number0, big.NewInt(0).Exp(multi, shiftDecimal, nil))

		resultNum0 = modifiedNumber0
		resultNum1 = number1
		return
	} else if decimal0 > decimal1 {
		shiftDecimal := big.NewInt(int64(decimal0 - decimal1))
		multi := big.NewInt(10)
		modifiedNumber1 := big.NewInt(0).Mul(number1, big.NewInt(0).Exp(multi, shiftDecimal, nil))

		resultNum0 = number0
		resultNum1 = modifiedNumber1
		return
	} else {
		resultNum0 = number0
		resultNum1 = number1
		return
	}
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
			fmt.Println("Tx", vLog.TxHash.String())
			println("FromAmt", logExchangedRawData[2].(*big.Int).String(), "ToAmt", logExchangedRawData[3].(*big.Int).String())

		case "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822":
			fmt.Println("Tx", vLog.TxHash.String())
			parsed, err := abi.JSON(strings.NewReader(uniPair.UniPairABI))
			logExchangedRawData, err := parsed.Unpack("Swap", vLog.Data)
			if err != nil {
				break
			}
			println(logExchangedRawData[0].(*big.Int).String(),
				logExchangedRawData[1].(*big.Int).String(),
				logExchangedRawData[2].(*big.Int).String(), logExchangedRawData[3].(*big.Int).String())
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
