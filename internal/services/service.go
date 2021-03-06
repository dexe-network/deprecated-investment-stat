package services

import (
	"context"
	"dex-trades-parser/internal/services/routes"
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
	Info          *service_routes.InfoRoutes
	Pools         *service_routes.PoolsRoutes
	PoolTransfers *service_routes.PoolTransfersRoutes
	Trades        *service_routes.TradesRoutes
	Nonce         *service_routes.NonceRoutes
	User          *service_routes.UserRoutes
	Trader        *service_routes.TraderRoutes
}

func InitServices(ctx context.Context, cancel func(), log *zap.Logger, st *storage.Storage, r *runway.Runway) *Services {
	routesContext := service_routes.NewRoutesContext(ctx, cancel, log, st)
	return &Services{
		Subscriber: subscriber.NewSubscriber(ctx, cancel, log, st, r.ETH(), r.Parser()),
		Routes: &Routes{
			Info:          &service_routes.InfoRoutes{Context: routesContext},
			Pools:         &service_routes.PoolsRoutes{Context: routesContext},
			PoolTransfers: &service_routes.PoolTransfersRoutes{Context: routesContext},
			Trades:        &service_routes.TradesRoutes{Context: routesContext},
			Nonce:         &service_routes.NonceRoutes{Context: routesContext},
			User:          &service_routes.UserRoutes{Context: routesContext},
			Trader:        &service_routes.TraderRoutes{Context: routesContext},
		},
	}
}
