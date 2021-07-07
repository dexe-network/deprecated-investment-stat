package router

import (
	"context"
	"dex-trades-parser/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
)

func InitRouter(
	e *gin.Engine,
	ctx context.Context,
	log *zap.Logger,
	services *services.Services,
) {
	Register(e)
	e.POST("/subscription", services.Subscriber.CreateSomething)
	e.GET("/subscription", services.Subscriber.GetSomething)
	e.PATCH("/subscription", services.Subscriber.UpdateSomething)
	e.DELETE("/subscription", services.Subscriber.DeleteSomething)

	// Pools Routes
	e.GET(PoolsPrefix+"/", services.Routes.Pool.GetAll)

	return
}

func serveError(w http.ResponseWriter, status int, txt string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Go-Pprof", "1")
	w.Header().Del("Content-Disposition")
	w.WriteHeader(status)
	fmt.Fprintln(w, txt)
}

func durationExceedsWriteTimeout(r *http.Request, seconds float64) bool {
	srv, ok := r.Context().Value(http.ServerContextKey).(*http.Server)
	return ok && srv.WriteTimeout != 0 && seconds >= srv.WriteTimeout.Seconds()
}
