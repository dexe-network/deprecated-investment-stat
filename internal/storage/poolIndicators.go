package storage

import (
	"dex-trades-parser/internal/models"
	"go.uber.org/zap"
	"sync"
	"time"
)

type PoolIndicatorsStorage struct {
	*Storage
	table string
	cache []*models.PoolIndicators
	mu    *sync.Mutex
}

func NewPoolIndicatorsStorage(st *Storage) *PoolIndicatorsStorage {
	s := &PoolIndicatorsStorage{
		Storage: st,
		table:   "poolIndicators",
		mu:      new(sync.Mutex),
		cache:   []*models.PoolIndicators{},
	}

	go s.poolIndicatorsWorker()

	return s
}

func (s *PoolIndicatorsStorage) poolIndicatorsWorker() {
	for {
		if len(s.cache) > 0 {
			s.storePoolIndicatorsToBD()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *PoolIndicatorsStorage) storePoolIndicatorsToBD() {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := s.DB.Create(&s.cache)
	s.cache = nil
	if result.Error != nil {
		s.Log.Error("Can't save poolIndicators to DB", zap.Error(result.Error))
	}
}

func (s *PoolIndicatorsStorage) Save(poolIndicators *models.PoolIndicators) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = append(s.cache, poolIndicators)
	return
}
