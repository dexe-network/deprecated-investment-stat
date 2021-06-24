package storage

import (
	"context"
	"dex-trades-parser/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	Pool Pool
}

func NewRepo(st *Storage) Repo {
	return Repo{
		EthTrade: NewEthTradeStorage(st),
		Pool: NewPoolStorage(st),
	}
}

type EthTrade interface {
	Save(ethTrade *models.EthTrade) (err error)
}

type Pool interface {
	Save(pool *models.Pool) (err error)
}
