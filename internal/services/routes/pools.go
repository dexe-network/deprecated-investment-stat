package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PoolsRoutes struct {
	Context *RoutesContext
}

func (p *PoolsRoutes) GetAll(c *gin.Context) {
	var pools []models.Pool
	if err := p.Context.st.DB.Find(&pools).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: pools})

}