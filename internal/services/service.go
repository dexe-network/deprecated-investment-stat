package services

import (
	"context"
	services_routes_pools "dex-trades-parser/internal/services/routes"
	"dex-trades-parser/internal/storage"
	"dex-trades-parser/internal/subscriber"
	"dex-trades-parser/pkg/runway"
	"go.uber.org/zap"
)

type Services struct {
	Subscriber *subscriber.Subscriber
	Routes     *Routes
}

type Routes struct {
	Pool *services_routes_pools.PoolService
}

func InitServices(ctx context.Context, cancel func(), log *zap.Logger, st *storage.Storage, r *runway.Runway) *Services {
	return &Services{
		Subscriber: subscriber.NewSubscriber(ctx, cancel, log, st, r.ETH(), r.Parser()),
		Routes:     &Routes{Pool: services_routes_pools.NewPoolService(ctx, cancel, log, st)},
	}
}
