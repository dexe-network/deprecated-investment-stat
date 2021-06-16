package pkg

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func GracefulServer(log *zap.Logger, ctx context.Context, cancel func(), srv *http.Server, errCh chan error) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		cancel()
		panic(err)
	case <-sigs:
		cancel()
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Error("Shutdown with error", zap.Error(err))
		} else {
			log.Info("Correct shutdown")
		}
	}
}
