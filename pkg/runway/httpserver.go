package runway

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HttpServerConfig struct {
	Router *gin.Engine
}

func (r *Runway) HttpServer(cfg HttpServerConfig) (server *http.Server, errCh chan error) {
	if !r.flag.IsApp() {
		r.log.Fatal("runway: required app flags")
	}

	srv := &http.Server{
		Addr:           viper.GetString("app-addr"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   31 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1Mb
	}

	srv.Handler = cfg.Router

	errCh = make(chan error, 1)

	go func() {
		r.log.Debug("running api server on %s", zap.String("app-addr", viper.GetString("app-addr")))
		errCh <- srv.ListenAndServe()
	}()

	return srv, errCh
}
