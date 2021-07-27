package router

import (
	"context"
	_ "dex-trades-parser/docs" // docs is generated by Swag CLI, you have to import it.
	"dex-trades-parser/internal/services"
	"dex-trades-parser/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
)

const (
	PoolsPrefix         = "/pools"
	PoolTransfersPrefix = "/pool-transfers"
	InfoPrefix          = "/info"
	TradesPrefix        = "/trades"
	NoncePrefix         = "/nonce"
	UserPrefix          = "/user"
)

func InitRouter(
	e *gin.Engine,
	ctx context.Context,
	log *zap.Logger,
	services *services.Services,
	st *storage.Storage,
) {
	Register(e)

	// Pools Routes
	e.GET(PoolsPrefix+"/", services.Routes.Pools.GetAll)
	e.GET(PoolsPrefix+"/:wallet", services.Routes.Pools.GetPoolsByWallet)

	// Info Route
	e.GET(InfoPrefix+"/global-token-whitelist", services.Routes.Info.GetGlobalTokenWhitelist)

	// Pool Transfers Route
	e.GET(PoolTransfersPrefix+"/withdrawals/:wallet", services.Routes.PoolTransfers.GetWithdrawalsByWallet)
	e.GET(PoolTransfersPrefix+"/deposits/:wallet", services.Routes.PoolTransfers.GetDepositsByWallet)

	// Trades
	e.GET(TradesPrefix+"/:traderPool", services.Routes.Trades.GetTradesByPoolAdr)

	// Nonce
	e.GET(NoncePrefix+"/:wallet", services.Routes.Nonce.GetNonce)

	// User
	//e.GET(UserPrefix+"/:wallet", services.Routes.User.GetUserInfo)

	// Required [SIGN]
	e.PUT(UserPrefix+"/:wallet/avatar", СheckAuthSign(st), services.Routes.User.PutAvatarUpdate)
	e.POST(UserPrefix+"/signup", СheckAuthSign(st), services.Routes.User.PostSignUp)
	//

	// Swagger
	url := ginSwagger.URL("http://" + viper.GetString("app-addr") + "/swagger/doc.json") // The url pointing to API definition
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
