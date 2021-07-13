package service_routes

import (
	"context"
	"dex-trades-parser/internal/storage"
	"go.uber.org/zap"
)

type RoutesContext struct {
	ctx    context.Context
	cancel func()
	log    *zap.Logger
	st     *storage.Storage
}

func NewRoutesContext(
	ctx context.Context,
	cancel func(),
	log *zap.Logger,
	st *storage.Storage,
) (s *RoutesContext) {
	routesService := &RoutesContext{
		ctx:    ctx,
		cancel: cancel,
		log:    log,
		st:     st,
	}
	return routesService
}
