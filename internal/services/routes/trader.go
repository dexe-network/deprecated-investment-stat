package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TraderRoutes struct {
	Context *RoutesContext
}

type TraderInfoResponse struct {
	Fund int64 `json:"fund"`
	Copiers float64 `json:"copiers"`
}

// @Description Get Trader Info
// @Summary Get Trader Info
// @Tags Trader
// @Accept  json
// @Produce  json
// @Param traderWallet path string true "Trader wallet address"
// @Success 200 {object} response.S{data=TraderInfoResponse}
// @Failure 400 {object} response.E
// @Router /trader/{traderWallet} [get]
func (p *TraderRoutes) GetTraderInfo(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("traderWallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid traderWallet address",
		})
		return
	}

	result := &TraderInfoResponse{}

	traderWallet := c.Param("traderWallet")
	var allTraderPools []models.Pool
	if err := p.Context.st.DB.Find(&allTraderPools, "LOWER(\"creatorAdr\") = LOWER(?)", traderWallet).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "Pools DB request error",
		})
		return
	}

	if len(allTraderPools) <= 0 {
		response.Error(c, http.StatusAccepted, response.E{
			Code:    response.InvalidRequest,
			Message: "Trader pools not exist",
		})
		return
	}

	poolsAdrList := []string{}
	for _, pool := range allTraderPools {
		poolsAdrList = append(poolsAdrList, pool.PoolAdr)
	}

	var investorsCount int64
	if err := p.Context.st.DB.Model(&models.PoolTransfer{}).Distinct("\"wallet\"").
		Where("\"poolAdr\" IN ?", poolsAdrList).Count(&investorsCount).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "Transfers DB request error",
		})
		return
	}
	result.Fund = investorsCount

	var investorsLast24hCount int64
	if err := p.Context.st.DB.Model(&models.PoolTransfer{}).Distinct("\"wallet\"").
		Where("\"poolAdr\" IN ? AND \"createdAt\" >= ?", poolsAdrList, time.Now().AddDate(0, 0, -1)).Count(&investorsLast24hCount).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "Transfers DB request error",
		})
		return
	}
	result.Copiers = float64(investorsCount) / float64(investorsLast24hCount) * 100

	response.Success(c, http.StatusOK, response.S{Data: result})

}
