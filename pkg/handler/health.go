package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dex-trades-parser/pkg/response"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Check(c *gin.Context) {
	response.Success(c, http.StatusOK, response.S{})
}
