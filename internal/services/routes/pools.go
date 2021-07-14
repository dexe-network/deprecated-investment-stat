package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PoolsRoutes struct {
	Context *RoutesContext
}

// @Description Get All Trader Pools
// @Summary Get All Trader Pools
// @Tags Pools
// @Accept  json
// @Produce  json
// @Success 200 {object} response.S{data=[]models.Pool}
// @Failure 400 {object} response.E
// @Router /pools [get]
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

// @Description Get Pools By Creator Wallet
// @Summary Get Pools By Creator Wallet
// @Tags Pools
// @Accept  json
// @Produce  json
// @Param   wallet path string true "Pool creator wallet"
// @Success 200 {object} response.S{data=[]models.Pool}
// @Failure 400 {object} response.E
// @Router /pools/{wallet} [get]
func (p *PoolsRoutes) GetPoolsByWallet(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("wallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}

	var pools []models.Pool
	if err := p.Context.st.DB.Find(&pools, "\"creatorAdr\" = ?", c.Param("wallet")).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: pools})

}
