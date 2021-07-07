package storage

import (
	"dex-trades-parser/internal/models"
	"go.uber.org/zap"
	"sync"
	"time"
)

type TradeStorage struct {
	*Storage
	table string
	cache []*models.Trade
	mu    *sync.Mutex
}

func NewTradeStorage(st *Storage) *TradeStorage {
	s := &TradeStorage{
		Storage: st,
		table:   "trades",
		mu:      new(sync.Mutex),
		cache:   []*models.Trade{},
	}

	go s.tradeWorker()

	return s
}

func (s *TradeStorage) tradeWorker() {
	for {
		if len(s.cache) > 0 {
			s.storeTradeToBD()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *TradeStorage) storeTradeToBD() {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := s.DB.Create(&s.cache)
	s.cache = nil
	if result.Error != nil {
		s.Log.Error("Can't save trade to DB", zap.Error(result.Error))
	}
}

func (s *TradeStorage) Save(trade *models.Trade) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = append(s.cache, trade)
	return
}
