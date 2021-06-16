package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Root struct {
	appName []byte
}

func NewRoot(appName string) *Root {
	return &Root{
		appName: []byte(appName),
	}
}

func (h *Root) Index(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", h.appName)
}
