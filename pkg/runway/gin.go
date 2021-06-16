package runway

import (
	"dex-trades-parser/pkg/handler"
	"dex-trades-parser/pkg/middleware/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GinConfig struct {
	Handlers []HandlerSetter
}

type HandlerSetter interface {
	SetHandler(e *gin.Engine)
}

func (r *Runway) Gin(config GinConfig) (e *gin.Engine) {
	e = gin.New()

	e.Use(gin.Recovery())

	if r.flag.IsCORS() {
		e.Use(api.CORSMiddleware(&api.CORSConfig{
			AllowCredentials: viper.GetBool("cors-allow-credentials"),
			AllowedHeaders:   viper.GetStringSlice("cors-allowed-headers"),
			ExposedHeaders:   viper.GetStringSlice("cors-exposed-headers"),
			AllowedMethods:   viper.GetStringSlice("cors-allowed-methods"),
			AllowedOrigins:   viper.GetStringSlice("cors-allowed-origins"),
		}))
	}

	for _, h := range config.Handlers {
		h.SetHandler(e)
	}

	if viper.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
	}

	return
}

type RootHandler struct {
	AppName string
}

func (h RootHandler) SetHandler(e *gin.Engine) {
	appName := h.AppName
	if appName == "" {
		appName = viper.GetString("app-name")
	}

	e.GET("", handler.NewRoot(appName).Index)
}

type HealthHandler struct{}

func (h HealthHandler) SetHandler(e *gin.Engine) {
	e.GET("/health", handler.NewHealth().Check)
}
