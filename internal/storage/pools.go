package storage

import (
	"dex-trades-parser/internal/models"
	"go.uber.org/zap"
	"sync"
	"time"
)

type PoolStorage struct {
	*Storage
	table string
	cache []*models.Pool
	mu    *sync.Mutex
}

func NewPoolStorage(st *Storage) *PoolStorage {
	s := &PoolStorage{
		Storage: st,
		table:   "pools",
		mu:      new(sync.Mutex),
		cache:   []*models.Pool{},
	}

	go s.poolWorker()

	return s
}

func (s *PoolStorage) poolWorker() {
	for {
		if len(s.cache) > 0 {
			s.storePoolToBD()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *PoolStorage) storePoolToBD() {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := s.DB.Create(&s.cache)
	s.cache = nil
	if result.Error != nil {
		s.Log.Error("Can't save pool to DB", zap.Error(result.Error))
	}
}

func (s *PoolStorage) Save(pool *models.Pool) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = append(s.cache, pool)
	return
}
