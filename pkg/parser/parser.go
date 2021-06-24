package parser

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"strings"
)

const ETHAddress = "0x0000000000000000000000000000000000000000"

type Parser struct {
	log            *zap.Logger
	basicAddresses BasicAddresses
	client         *ethclient.Client
	Abis           Abis
}

type BasicAddresses struct {
	traderPoolFactoryUpgradeable common.Address
}

type Abis struct {
	erc20             abi.ABI
	traderPoolFactory abi.ABI
	traderPool        abi.ABI
}

func NewParser(log *zap.Logger, client *ethclient.Client, factoryAddress string) (p *Parser, err error) {
	erc20, err := loadAbiByName("ERC20", log)
	traderPoolFactory, err := loadAbiByName("TraderPoolFactoryUpgradeable", log)
	traderPool, err := loadAbiByName("TraderPoolUpgradeable", log)

	p = &Parser{
		log:            log,
		basicAddresses: BasicAddresses{traderPoolFactoryUpgradeable: common.HexToAddress(factoryAddress)},
		Abis: Abis{
			erc20,
			traderPoolFactory,
			traderPool,
		},
		client: client,
	}

	return
}

func loadAbiByName(name string, log *zap.Logger) (parseAbi abi.ABI, err error) {
	rawBytes, err := ioutil.ReadFile("pkg/abi/" + name + ".json")
	if err != nil {
		log.Debug("ioutil.ReadFile", zap.Error(err))
		return
	}

	parseAbi, err = abi.JSON(strings.NewReader(string(rawBytes)))
	if err != nil {
		log.Debug("abi.JSON", zap.Error(err))
		return
	}

	return
}

func (p *Parser) FactoryAddress() string {
	return p.basicAddresses.traderPoolFactoryUpgradeable.Hex()
}

type ParsedTx struct {
	TokenA    common.Address
	TokenB    common.Address
	Tx        string
	Protocol  string
	AmountIn  *big.Int
	AmountOut *big.Int
	Wallet    common.Address
}

type ParsedFactoryTx struct {
	CreatorAdr            common.Address
	BasicTokenAdr         common.Address
	TotalSupply           string
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
	method, err := p.Abis.traderPoolFactory.MethodById(t.Data()[:4])
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
		pTx.TotalSupply = data["_totalSupply"].(*big.Int).String()
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
		pTx.PoolAdr = receipt.Logs[len(receipt.Logs) - 1].Address.String()

	default:
		fmt.Println("Unknown metod : " + method.Name)
		err = errors.New("Unknown metod : " + method.Name)
	}

	return
}

//func (p *Parser) Parse(t types.Transaction) (pTx ParsedTx, err error) {
//
//	receipt, err := p.client.TransactionReceipt(context.Background(), t.Hash())
//	if err != nil {
//		return
//	} else {
//		if receipt.Status != 1 {
//			err = errors.New("transaction status fail : " + t.Hash().String())
//			return
//		}
//	}
//
//	if len(t.Data()) < 5 {
//		err = errors.New("transaction data to small")
//		p.log.Debug("Tx Data len < 5", zap.Error(err), zap.String("Tx", t.Hash().String()))
//		return
//	}
//
//	data := make(map[string]interface{})
//	method, err := p.abi.MethodById(t.Data()[:4])
//	if err != nil {
//		p.log.Debug("Tx Data len < 5", zap.Error(err), zap.String("Tx", t.Hash().String()))
//		return
//	}
//	err = method.Inputs.UnpackIntoMap(data, t.Data()[4:])
//	if err != nil {
//		p.log.Debug("UnpackIntoMap", zap.Error(err), zap.String("Tx", t.Hash().String()))
//		return
//	}
//
//	pTx.Tx = t.Hash().String()
//	pTx.Protocol = p.protocol
//
//	switch method.Name {
//
//	case "swapExactTokensForTokens":
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForTokens cent get To")
//			return
//		}
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForTokens cent get Path")
//			return
//		}
//		AmountIn, ok := data["amountIn"].(*big.Int)
//		if !ok {
//			err = errors.New("swapExactTokensForTokens cent get AmountIn")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = Path[0]
//		pTx.TokenB = Path[len(Path)-1]
//		pTx.AmountIn = AmountIn
//		pTx.AmountOut, err = p.getOutAmount(receipt, pTx.TokenB.String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//	case "swapExactTokensForTokensSupportingFeeOnTransferTokens":
//		AmountIn, ok := data["amountIn"].(*big.Int)
//		if !ok {
//			err = errors.New("swapExactTokensForTokensSupportingFeeOnTransferTokens cent get AmountIn")
//			return
//		}
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForTokensSupportingFeeOnTransferTokens cent get Path")
//			return
//		}
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForTokensSupportingFeeOnTransferTokens cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = Path[0]
//		pTx.TokenB = Path[len(Path)-1]
//		pTx.AmountIn = AmountIn
//		pTx.AmountOut, err = p.getOutAmount(receipt, pTx.TokenB.String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//	case "swapTokensForExactTokens":
//		AmountInMax, ok := data["amountInMax"].(*big.Int)
//		if !ok {
//			err = errors.New("swapTokensForExactTokens cent get AmountInMax")
//			return
//		}
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapTokensForExactTokens cent get Path")
//			return
//		}
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapTokensForExactTokens cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = Path[0]
//		pTx.TokenB = Path[len(Path)-1]
//		pTx.AmountIn = AmountInMax
//		pTx.AmountOut, err = p.getOutAmount(receipt, pTx.TokenB.String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//
//	case "swapETHForExactTokens":
//		Amount := t.Value()
//		if Amount.Cmp(big.NewInt(0)) <= 0 {
//			err = errors.New("swapETHForExactTokens cent get Amount")
//			return
//		}
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapETHForExactTokens cent get Path")
//			return
//		}
//
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapETHForExactTokens cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = common.HexToAddress(ETHAddress)
//		pTx.TokenB = Path[len(Path)-1]
//		pTx.AmountIn = Amount
//		pTx.AmountOut, err = p.getOutAmount(receipt, pTx.TokenB.String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//	case "swapExactETHForTokens":
//		Amount := t.Value()
//		if Amount.Cmp(big.NewInt(0)) <= 0 {
//			err = errors.New("swapExactETHForTokens cent get Amount")
//			return
//		}
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapExactETHForTokens cent get Path")
//			return
//		}
//
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapExactETHForTokens cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = common.HexToAddress(ETHAddress)
//		pTx.TokenB = Path[len(Path)-1]
//		pTx.AmountIn = Amount
//		pTx.AmountOut, err = p.getOutAmount(receipt, pTx.TokenB.String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//	case "swapExactETHForTokensSupportingFeeOnTransferTokens":
//		Amount := t.Value()
//		if Amount.Cmp(big.NewInt(0)) <= 0 {
//			err = errors.New("swapExactETHForTokensSupportingFeeOnTransferTokens cent get Amount")
//			return
//		}
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapExactETHForTokensSupportingFeeOnTransferTokens cent get Path")
//			return
//		}
//
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapExactETHForTokensSupportingFeeOnTransferTokens cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = common.HexToAddress(ETHAddress)
//		pTx.TokenB = Path[len(Path)-1]
//		pTx.AmountIn = Amount
//		pTx.AmountOut, err = p.getOutAmount(receipt, pTx.TokenB.String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//
//	case "swapExactTokensForETH":
//		AmountIn, ok := data["amountIn"].(*big.Int)
//		if !ok {
//			err = errors.New("swapExactTokensForETH cent get AmountIn")
//			return
//		}
//
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForETH cent get Path")
//			return
//		}
//
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForETH cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = Path[0]
//		pTx.TokenB = common.HexToAddress(ETHAddress)
//		pTx.AmountIn = AmountIn
//		pTx.AmountOut, err = p.getOutAmount(receipt, Path[len(Path)-1].String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//	case "swapTokensForExactETH":
//		AmountInMax, ok := data["amountInMax"].(*big.Int)
//		if !ok {
//			err = errors.New("swapTokensForExactETH cent get AmountInMax")
//			return
//		}
//
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapTokensForExactETH cent get Path")
//			return
//		}
//
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapTokensForExactETH cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = Path[0]
//		pTx.TokenB = common.HexToAddress(ETHAddress)
//		pTx.AmountIn = AmountInMax
//		pTx.AmountOut, err = p.getOutAmount(receipt, Path[len(Path)-1].String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//	case "swapExactTokensForETHSupportingFeeOnTransferTokens":
//		AmountIn, ok := data["amountIn"].(*big.Int)
//		if !ok {
//			err = errors.New("swapExactTokensForETHSupportingFeeOnTransferTokens cent get AmountIn")
//			return
//		}
//
//		Path, ok := data["path"].([]common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForETHSupportingFeeOnTransferTokens cent get Path")
//			return
//		}
//
//		To, ok := data["to"].(common.Address)
//		if !ok {
//			err = errors.New("swapExactTokensForETHSupportingFeeOnTransferTokens cent get To")
//			return
//		}
//
//		pTx.Wallet = To
//		pTx.TokenA = Path[0]
//		pTx.TokenB = common.HexToAddress(ETHAddress)
//		pTx.AmountIn = AmountIn
//		pTx.AmountOut, err = p.getOutAmount(receipt, Path[len(Path)-1].String(), pTx.TokenB.String(), pTx.Wallet.String())
//		return
//	}
//
//	err = errors.New("Unknown pancake swap metod : " + method.Name)
//	return
//}
//
//func (p *Parser) getOutAmount(r *types.Receipt, tokenB string, fTokenB string, wallet string) (OutAmount *big.Int, err error) {
//
//	withdrawalSig := crypto.Keccak256Hash([]byte("Withdrawal(address,uint256)"))
//	transferSig := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
//
//	for _, log := range r.Logs {
//		if log.Topics[0].Hex() == withdrawalSig.Hex() && len(log.Topics) > 1 && fTokenB == ETHAddress {
//
//			if common.HexToAddress(log.Topics[1].Hex()).String() == p.routerAddress.Hex() && log.Address.String() == tokenB {
//				OutAmount = new(big.Int).SetBytes(log.Data)
//				return OutAmount, nil
//			}
//		}
//
//		if log.Topics[0].Hex() == transferSig.Hex() && len(log.Topics) > 2 && tokenB != ETHAddress {
//			if common.HexToAddress(log.Topics[2].Hex()).String() == wallet && log.Address.String() == tokenB {
//				OutAmount = new(big.Int).SetBytes(log.Data)
//				return OutAmount, nil
//			}
//		}
//	}
//
//	err = errors.New("get OutAmount fail tx: " + r.TxHash.Hex())
//	return
//
//}
