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
	mu *sync.Mutex
}

func NewEthTradeStorage(st *Storage) *EthTradeStorage {
	s := &EthTradeStorage{
		Storage: st,
		table:   "eth_trades",
		mu: new(sync.Mutex),
		cache: []*models.EthTrade{},
	}


	go s.worker()

	return s
}

func (s *EthTradeStorage) worker(){
	for {
		if len(s.cache) > 0 {
			s.storeToBD()
		}
		time.Sleep(100 * time.Millisecond)
	}
}


func (s *EthTradeStorage) storeToBD(){
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


	//s.DB.Exec(`
	//	INSERT INTO "eth_trades" (
	//		"tokenA",
	//		"tokenB",
	//		"date",
	//		"blockNumber",
	//		"tx",
	//		"protocol",
	//		"priceIn",
	//		"priceOut",
	//		"amountOut",
	//		"amountIn",
	//		"wallet",
	//		"value"
	//	) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
	//		ethTrade.TokenA,			//tokenA   '0x0000000000000000000000000000000000000000',
	//		ethTrade.TokenB,			//tokenB   '0x15874d65e649880c2614e7a480cb7c9A55787FF6',
	//		ethTrade.Date.UTC(),		//date   '2021-06-02 19:40:51',
	//		ethTrade.BlockNumber,		//blockNumber   12556233,
	//		ethTrade.Tx,				//tx   '0xe9ff6df4967c6e0c1901d3a01869de839499f19d10f92dce3d125eaf341f109b',
	//		ethTrade.Protocol,			//protocol   'uniswapV2',
	//		ethTrade.PriceIn.String(),	//priceIn   '0.0000000001591351',
	//		ethTrade.PriceOut.String(),	//priceOut   '6283968648.404942079165463',
	//		ethTrade.AmountOut,			//amountOut   '25135874593619768316661852',
	//		ethTrade.AmountIn,			//amountIn   '4000000000000000',
	//		ethTrade.Wallet,			//wallet   '0xD13751b33363CEB61360d8154a0EC6fc8Ee2bA16',
	//		ethTrade.Value,				//value   ''
	//)



	//return s.DB.
	//	Table(s.table).
	//	Create(&ethTrade).
	//	Debug().
	//	Error
	return
}
