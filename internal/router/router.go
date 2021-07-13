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

const (
	PoolsPrefix = "/pools"
	PoolTransfersPrefix = "/pool-transfers"
	InfoPrefix  = "/info"
)

func InitRouter(
	e *gin.Engine,
	ctx context.Context,
	log *zap.Logger,
	services *services.Services,
) {
	Register(e)

	// Pools Routes
	e.GET(PoolsPrefix+"/", services.Routes.Pools.GetAll)

	// Info Route
	e.GET(InfoPrefix+"/global-token-whitelist", services.Routes.Info.GetGlobalTokenWhitelist)

	// Pool Transfers Route
	e.GET(PoolTransfersPrefix+"/withdrawals/:wallet", services.Routes.PoolTransfers.GetWithdrawalsByWallet)
	e.GET(PoolTransfersPrefix+"/deposits/:wallet", services.Routes.PoolTransfers.GetDepositsByWallet)

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
