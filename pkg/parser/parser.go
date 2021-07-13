package parser

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
	exchangeTool                 common.Address
	paramKeeper                  common.Address
}

type Abis struct {
	Erc20             abi.ABI
	TraderPoolFactory abi.ABI
	TraderPool        abi.ABI
	ExchangeTool      abi.ABI
	ParamKeeper       abi.ABI
}

func NewParser(log *zap.Logger, client *ethclient.Client, factoryAddress string, exchangeToolAddress string, paramKeeperAddress string) (p *Parser, err error) {
	erc20, err := loadAbiByName("ERC20", log)
	traderPoolFactory, err := loadAbiByName("TraderPoolFactoryUpgradeable", log)
	traderPool, err := loadAbiByName("TraderPoolUpgradeable", log)
	exchangeTool, err := loadAbiByName("UniswapExchangeTool", log)
	paramKeeper, err := loadAbiByName("ParamKeeper", log)

	p = &Parser{
		log: log,
		basicAddresses: BasicAddresses{
			traderPoolFactoryUpgradeable: common.HexToAddress(factoryAddress),
			exchangeTool:                 common.HexToAddress(exchangeToolAddress),
			paramKeeper:                  common.HexToAddress(paramKeeperAddress),
		},
		Abis: Abis{
			erc20,
			traderPoolFactory,
			traderPool,
			exchangeTool,
			paramKeeper,
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

func (p *Parser) ExchangeToolAddress() string {
	return p.basicAddresses.exchangeTool.Hex()
}

func (p *Parser) ParamKeeperAddress() string {
	return p.basicAddresses.paramKeeper.Hex()
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
