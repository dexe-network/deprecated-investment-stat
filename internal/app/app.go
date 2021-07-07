package app

import (
	"context"
	"dex-trades-parser/internal/router"
	"dex-trades-parser/internal/services"
	"dex-trades-parser/internal/storage"
	"dex-trades-parser/pkg"
	"dex-trades-parser/pkg/flag"
	"dex-trades-parser/pkg/runway"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"
	"time"
)

func Run(ctx context.Context, cancel func(), fl *flag.Flag) {

	r := runway.NewRunway(fl)
	log := r.Log().WithOptions(zap.AddStacktrace(zap.ErrorLevel))
	st := storage.NewStorage(log, r.DB())
	logHook, _ := zap.NewProduction()

	e := r.Gin(runway.GinConfig{Handlers: []runway.HandlerSetter{
		runway.RootHandler{},
		runway.HealthHandler{},
	}})

	e.Use(ginzap.Ginzap(logHook, time.RFC3339, true))
	e.Use(ginzap.RecoveryWithZap(logHook, true))

	servicesInst := services.InitServices(ctx, cancel, log, st, r)

	router.InitRouter(e, ctx, log, servicesInst)
	srv, errCh := r.HttpServer(runway.HttpServerConfig{
		Router: e,
	})

	go pkg.GracefulServer(log, ctx, cancel, srv, errCh)

	fmt.Println("START!!!")
	servicesInst.Subscriber.Run()

	<-ctx.Done()
}
