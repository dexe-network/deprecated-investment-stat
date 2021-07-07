package storage

import (
	"dex-trades-parser/internal/models"
	"fmt"
	"strings"
	"sync"
	"time"
)

type EthTradeStorage struct {
	*Storage
	table string
	cache []*models.EthTrade
	mu    *sync.Mutex
}

func NewEthTradeStorage(st *Storage) *EthTradeStorage {
	s := &EthTradeStorage{
		Storage: st,
		table:   "eth_trades",
		mu:      new(sync.Mutex),
		cache:   []*models.EthTrade{},
	}

	go s.worker()

	return s
}

func (s *EthTradeStorage) worker() {
	for {
		if len(s.cache) > 0 {
			s.storeToBD()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *EthTradeStorage) storeToBD() {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := []string{}
	valuesArgs := []interface{}{}
	for _, ethTrade := range s.cache {
		values = append(values, "( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

		valuesArgs = append(valuesArgs, ethTrade.TokenA)
		valuesArgs = append(valuesArgs, ethTrade.TokenB)
		valuesArgs = append(valuesArgs, ethTrade.Date.UTC())
		valuesArgs = append(valuesArgs, ethTrade.BlockNumber)
		valuesArgs = append(valuesArgs, ethTrade.Tx)
		valuesArgs = append(valuesArgs, ethTrade.Protocol)
		valuesArgs = append(valuesArgs, ethTrade.PriceIn.String())
		valuesArgs = append(valuesArgs, ethTrade.PriceOut.String())
		valuesArgs = append(valuesArgs, ethTrade.AmountOut)
		valuesArgs = append(valuesArgs, ethTrade.AmountIn)
		valuesArgs = append(valuesArgs, ethTrade.Wallet)
		valuesArgs = append(valuesArgs, ethTrade.Value)
	}

	query := `INSERT INTO "eth_trades" (
			"tokenA",
			"tokenB",
			"date",
			"blockNumber",
			"tx",
			"protocol",
			"priceIn",
			"priceOut",
			"amountOut",
			"amountIn",
			"wallet",
			"value"
		) VALUES %s `
	query = fmt.Sprintf(query, strings.Join(values, ","))
	fmt.Println("query:", query)
	s.DB.Exec(query, valuesArgs...)
	//log.Fatal("DONE")
}

func (s *EthTradeStorage) Save(ethTrade *models.EthTrade) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = append(s.cache, ethTrade)
	return
}
