package storage

import (
	"dex-trades-parser/internal/models"
	"go.uber.org/zap"
	"sync"
	"time"
)

type PoolTransfersStorage struct {
	*Storage
	table string
	cache []*models.PoolTransfer
	mu    *sync.Mutex
}

func NewPoolTransfersStorage(st *Storage) *PoolTransfersStorage {
	s := &PoolTransfersStorage{
		Storage: st,
		table:   "poolTransfers",
		mu:      new(sync.Mutex),
		cache:   []*models.PoolTransfer{},
	}

	go s.poolTransfersWorker()

	return s
}

func (s *PoolTransfersStorage) poolTransfersWorker() {
	for {
		if len(s.cache) > 0 {
			s.storePoolTransfersToBD()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *PoolTransfersStorage) storePoolTransfersToBD() {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := s.DB.Create(&s.cache)
	s.cache = nil
	if result.Error != nil {
		s.Log.Error("Can't save poolTransfer to DB", zap.Error(result.Error))
	}
}

func (s *PoolTransfersStorage) Save(poolTransfer *models.PoolTransfer) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = append(s.cache, poolTransfer)
	return
}
