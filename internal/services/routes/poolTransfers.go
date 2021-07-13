package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PoolTransfersRoutes struct {
	Context *RoutesContext
}

func (p *PoolTransfersRoutes) GetWithdrawalsByWallet(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("wallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}

	var transfers []models.PoolTransfer

	if err := p.Context.st.DB.Find(
		&transfers, "type = ? AND LOWER(wallet) = LOWER(?)", "deposit", c.Param("wallet"),
	).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: transfers})

}

func (p *PoolTransfersRoutes) GetDepositsByWallet(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("wallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}

	var transfers []models.PoolTransfer

	if err := p.Context.st.DB.Find(
		&transfers, "type = ? AND LOWER(wallet) = LOWER(?)", "withdraw", c.Param("wallet"),
	).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: transfers})

}
