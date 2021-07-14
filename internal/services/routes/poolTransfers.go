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

// @Description Get Withdrawals By Wallet
// @Summary Get Withdrawals By Wallet
// @Tags PoolTransfers
// @Accept  json
// @Produce  json
// @Param   wallet path string true "Wallet Address"
// @Success 200 {object} response.S{data=[]models.PoolTransfer}
// @Failure 400 {object} response.E
// @Router /pool-transfers/withdrawals/{wallet} [get]
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

// @Description Get Deposits By Wallet
// @Summary Get Deposits By Wallet
// @Tags PoolTransfers
// @Accept  json
// @Produce  json
// @Param   wallet path string true "Wallet Address"
// @Success 200 {object} response.S{data=[]models.PoolTransfer}
// @Failure 400 {object} response.E
// @Router /pool-transfers/deposits/{wallet} [get]
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
