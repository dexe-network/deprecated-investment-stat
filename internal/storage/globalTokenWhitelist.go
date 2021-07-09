package storage

import (
	"dex-trades-parser/internal/models"
	"go.uber.org/zap"
	"sync"
	"time"
)

type GlobalTokenWhitelistStorage struct {
	*Storage
	table string
	cache []*models.GlobalTokenWhitelist
	mu    *sync.Mutex
}

func NewGlobalTokenWhitelistStorage(st *Storage) *GlobalTokenWhitelistStorage {
	s := &GlobalTokenWhitelistStorage{
		Storage: st,
		table:   "globalTokenWhitelist",
		mu:      new(sync.Mutex),
		cache:   []*models.GlobalTokenWhitelist{},
	}

	go s.globalTokenWhitelistWorker()

	return s
}

func (s *GlobalTokenWhitelistStorage) globalTokenWhitelistWorker() {
	for {
		if len(s.cache) > 0 {
			s.storeGlobalTokenWhitelistToBD()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *GlobalTokenWhitelistStorage) storeGlobalTokenWhitelistToBD() {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := s.DB.Create(&s.cache)
	s.cache = nil
	if result.Error != nil {
		s.Log.Error("Can't save GlobalTokenWhitelist to DB", zap.Error(result.Error))
	}
}

func (s *GlobalTokenWhitelistStorage) Save(item *models.GlobalTokenWhitelist) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = append(s.cache, item)
	return
}

func (s *GlobalTokenWhitelistStorage) Delete(address string) (err error) {
	err = s.DB.Where("address = ?", address).Delete(&models.GlobalTokenWhitelist{}).Error
	return
}
