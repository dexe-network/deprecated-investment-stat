package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AccountIDKey   = "accountID"
	RoleKey        = "role"
	AccessTokenKey = "accessToken"
)

func CORSMiddleware(cfg *CORSConfig) func(c *gin.Context) {
	return func(c *gin.Context) {
		if cfg.AllowCredentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		c.Header("Access-Control-Allow-Methods", strings.Join(cfg.AllowedMethods, ","))
		c.Header("Access-Control-Allow-Headers", strings.Join(cfg.AllowedHeaders, ","))
		c.Header("Access-Control-Expose-Headers", strings.Join(cfg.ExposedHeaders, ","))

		for _, v := range cfg.AllowedOrigins {
			if v == c.GetHeader("Origin") {
				c.Header("Access-Control-Allow-Origin", v)
			}
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
