package subscriber

import (
	"context"
	"dex-trades-parser/internal/contracts/erc20"
	"dex-trades-parser/internal/models"
	//"dex-trades-parser/internal/models"
	"dex-trades-parser/internal/storage"
	"dex-trades-parser/pkg/parser"
	"dex-trades-parser/pkg/response"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Subscriber struct {
	ctx           context.Context
	cancel        func()
	log           *zap.Logger
	st            *storage.Storage
	client        *ethclient.Client
	newHeadSub    ethereum.Subscription
	parser        *parser.Parser
	headChan      chan *types.Header
	blockNumber   int64
	blockNumberMu *sync.Mutex
}

func NewSubscriber(
	ctx context.Context,
	cancel func(),
	log *zap.Logger,
	st *storage.Storage,
	client *ethclient.Client,
	parser *parser.Parser,
) (s *Subscriber) {
	subscriber := &Subscriber{
		ctx:           ctx,
		cancel:        cancel,
		log:           log,
		st:            st,
		client:        client,
		parser:        parser,
		headChan:      make(chan *types.Header, 10000),
		blockNumber:   0,
		blockNumberMu: new(sync.Mutex),
	}
	subscriber.loadBlockNumberFromFile()
	return subscriber
}

type SubscribeRequest struct {
	Data int64 `json:"data"`
}

func (s *Subscriber) CreateSomething(c *gin.Context) {
	sr := SubscribeRequest{}
	if err := c.ShouldBind(&sr); err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	//Do Some thing here

	response.Success(c, http.StatusOK, response.S{})

}

func (s *Subscriber) GetSomething(c *gin.Context) {
	sr := SubscribeRequest{}

	if err := c.ShouldBind(&sr); err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	//Do Some thing here

	response.Success(c, http.StatusOK, response.S{Data: 34})

}

func (s *Subscriber) UpdateSomething(c *gin.Context) {
	sr := SubscribeRequest{}
	if err := c.ShouldBind(&sr); err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	//Do Some thing here

	response.Success(c, http.StatusOK, response.S{})

}

func (s *Subscriber) DeleteSomething(c *gin.Context) {
	sr := SubscribeRequest{}
	if err := c.ShouldBind(&sr); err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	//Do Some thing here

	response.Success(c, http.StatusOK, response.S{})

}

func (s *Subscriber) Run() {

	if s.blockNumber > 0 {
		s.loadHistory()
	}

	s.subscribe()
	go s.subscriptionHandler()

}

func (s *Subscriber) subscribe() {

	newHeadSub, err := s.client.SubscribeNewHead(context.Background(), s.headChan)
	if err != nil {
		log.Fatal(err)
	}
	s.newHeadSub = newHeadSub
}

func (s *Subscriber) subscriptionHandler() {

	for {
		select {
		case err := <-s.newHeadSub.Err():
			log.Fatal(err)
		case header := <-s.headChan:
			if header.Number.Int64() <= s.blockNumber {
				break
			}
			go s.handleHeader(*header)
		}
	}
}

func (s *Subscriber) loadHistory() {

	var lastBlockNumber int64
	lastBlockID, err := s.client.BlockNumber(context.Background())
	if err == nil {
		lastBlockNumber = int64(lastBlockID)
	}

	maxGoroutines := viper.GetInt64("max-parallel-blocks")
	guard := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	for i := s.blockNumber; i <= lastBlockNumber; i++ {
		guard <- struct{}{}
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			headerNumber := big.NewInt(i)

			Header, err := s.client.HeaderByNumber(context.Background(), headerNumber)
			if err != nil {
				log.Fatal(err)
			}

			s.handleHeader(*Header)
			<-guard
		}(i)

		lastBlockID, err := s.client.BlockNumber(context.Background())
		if err == nil {
			lastBlockNumber = int64(lastBlockID)
		}
	}
}

func (s *Subscriber) handleHeader(header types.Header) {

	trie := 0
	var block *types.Block
	for {
		b, err := s.client.BlockByHash(context.Background(), header.Hash())
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			trie++
			s.log.Debug(header.Hash().String() + "Cent get block by hash, will retry in 100 Milliseconds")
		} else {
			block = b
			break
		}
	}

	var wg sync.WaitGroup
	for _, t := range block.Transactions() {
		if t.To() == nil {
			continue
		}

		switch t.To().Hex() {
		// Catch Any ParamKeeper operations
		case s.parser.ParamKeeperAddress():
			fmt.Println("handleHeader", "ParamKeeper")
			wg.Add(1)
			go func(t types.Transaction, blockNumber int64) {
				defer wg.Done()
				s.paramKeeperTransactionProcessing(t, block.Number().Int64(), block.Time())
			}(*t, block.Number().Int64())
		// Catch Any Exchange Tool operations
		case s.parser.ExchangeToolAddress():
			fmt.Println("handleHeader", "ExchangeTool")
			wg.Add(1)
			go func(t types.Transaction, blockNumber int64) {
				defer wg.Done()
				s.exchangeToolTransactionProcessing(t, block.Number().Int64(), block.Time())
			}(*t, block.Number().Int64())
		// Catch Any trader Pool operations
		case s.isTraderPoolTx(t.To()):
			fmt.Println("handleHeader", "TraderPool")
			wg.Add(1)
			go func(t types.Transaction, blockNumber int64) {
				defer wg.Done()
				s.traderPoolTransactionProcessing(t, block.Number().Int64(), block.Time())
			}(*t, block.Number().Int64())

		// Catch New Traders Pool
		case s.parser.FactoryAddress():
			fmt.Println("handleHeader", "FactoryAddress")
			wg.Add(1)
			go func(t types.Transaction, blockNumber int64) {
				defer wg.Done()
				s.factoryTransactionProcessing(t, block.Number().Int64(), block.Time())
			}(*t, block.Number().Int64())
		}
	}
	wg.Wait()

	s.updateBlockNumber(header.Number.Int64())

	var lastBlockNumber int64
	lastBlockID, err := s.client.BlockNumber(context.Background())
	if err == nil {
		lastBlockNumber = int64(lastBlockID)
	}
	fmt.Println(time.Now().String(), header.Number, "/", lastBlockNumber)

	return
}

func (s *Subscriber) isTraderPoolTx(to *common.Address) (hexAddress string) {
	var pool models.Pool
	if err := s.st.DB.First(&pool, "\"poolAdr\" = ?", to.String()).Error; err != nil {
		s.log.Debug("isTraderPoolTx Db Request Error", zap.Error(err))
		return
	}
	hexAddress = pool.PoolAdr
	return
}

func (s *Subscriber) updateBlockNumber(bn int64) {
	s.blockNumberMu.Lock()
	defer s.blockNumberMu.Unlock()

	if bn < s.blockNumber {
		return
	}

	s.blockNumber = bn

	err := ioutil.WriteFile("block_number", []byte(strconv.FormatInt(bn, 10)), 0644)
	if err != nil {
		s.log.Error("ioutil.WriteFile", zap.Error(err))
	}
}

func (s *Subscriber) loadBlockNumberFromFile() {
	rawBytes, err := ioutil.ReadFile("block_number")
	if err != nil {
		s.log.Debug("ioutil.ReadFile", zap.Error(err))
		return
	}

	s.blockNumber, err = strconv.ParseInt(strings.TrimSpace(string(rawBytes)), 10, 64)
	if err != nil {
		s.log.Error("strconv.ParseInt", zap.Error(err))
		return
	}
}

func (s *Subscriber) getPrices(Tx string, TokenIn common.Address, TokenOut common.Address, in *big.Int, out *big.Int) (priceIn decimal.Decimal, priceOut decimal.Decimal, err error) {

	inDec, err := s.getTokenDecimals(TokenIn)
	if err != nil {
		s.log.Error("cent create erc20 instance: "+TokenIn.String(), zap.Error(err))
		return
	}
	outDec, err := s.getTokenDecimals(TokenOut)
	if err != nil {
		s.log.Error("cent create erc20 instance: "+TokenOut.String(), zap.Error(err))
		return
	}

	DecimalIn := decimal.NewFromBigInt(in, 0)
	DecimalOut := decimal.NewFromBigInt(out, 0)
	DecimalInDec := decimal.New(int64(inDec), 0)
	DecimalOutDec := decimal.New(int64(outDec), 0)

	if DecimalOut.IsZero() || DecimalIn.IsZero() {
		err = errors.New("DecimalOut IsZero")
		return
	}
	if DecimalIn.IsZero() {
		err = errors.New("DecimalIn IsZero")
		return
	}

	priceIn = DecimalIn.Div(DecimalOut).Mul(decimal.New(10, 0).Pow(DecimalOutDec.Sub(DecimalInDec)))
	priceOut = DecimalOut.Div(DecimalIn).Mul(decimal.New(10, 0).Pow(DecimalInDec.Sub(DecimalOutDec)))

	return
}

func (s *Subscriber) getTokenDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	if tokenAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		decimals = 18
		return
	}

	instance, err := erc20.NewErc20(tokenAddress, s.client)
	if err != nil {
		return
	}

	decimals, err = instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return
	}

	return
}
