package storage

import (
	"context"
	"dex-trades-parser/internal/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Storage struct {
	Ctx  context.Context
	Log  *zap.Logger
	DB   *gorm.DB
	Repo Repo
}

func NewStorage(
	log *zap.Logger,
	db *gorm.DB,
) *Storage {
	st := &Storage{
		Log: log,
		DB:  db,
	}

	st.Repo = NewRepo(st)

	return st
}

type Repo struct {
	EthTrade EthTrade
}

func NewRepo(st *Storage) Repo {
	return Repo{
		EthTrade: NewEthTradeStorage(st),
	}
}

type EthTrade interface {
	Save(ethTrade *models.EthTrade) (err error)
}
