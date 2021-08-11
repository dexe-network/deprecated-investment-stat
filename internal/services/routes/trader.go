package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type TraderRoutes struct {
	Context *RoutesContext
}

type PoolInfoResponse struct {
	Fund           int64   `json:"fund"`
	Copiers        float64 `json:"copiers"`
	Symbol         string  `json:"symbol"`
	BasicTokenAdr  string  `json:"basicTokenAdr"`
	CurrentPrice   float64 `json:"currentPrice"`
	PriceChange24H float64 `json:"priceChange24H"`
}

// @Description Get Trader/Pool info
// @Summary Get Trader/Pool info
// @Tags Trader
// @Accept  json
// @Produce  json
// @Param poolAddress path string true "Pool address"
// @Success 200 {object} response.S{data=PoolInfoResponse}
// @Failure 400 {object} response.E
// @Router /trader/{poolAddress}/info [get]
func (p *TraderRoutes) GetPoolInfo(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("poolAddress")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid pool Address",
		})
		return
	}

	result := &PoolInfoResponse{}

	////// Pool Data
	poolAddress := c.Param("poolAddress")
	var foundPool models.Pool
	if err := p.Context.st.DB.First(&foundPool, "LOWER(\"poolAdr\") = LOWER(?)", poolAddress).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "Pool not found",
		})
		return
	}
	result.Symbol = foundPool.Symbol
	result.BasicTokenAdr = foundPool.BasicTokenAdr
	//////

	////// Indicators Data
	var indicatorLast models.PoolIndicators
	if err := p.Context.st.DB.Order("date desc").First(&indicatorLast,
		"\"poolAdr\" = ?", foundPool.PoolAdr).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "Indicators DB request error",
		})
		return
	}

	totalCap, err := strconv.ParseFloat(indicatorLast.TotalCap, 64)
	totalSupply, err := strconv.ParseFloat(indicatorLast.TotalSupply, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "ParseFloat request error",
		})
		return
	}
	if totalCap <= 0 || totalSupply <= 0 {
		result.CurrentPrice = 0
	} else {
		result.CurrentPrice = totalCap / totalSupply
	}

	var indicatorsLast24h []models.PoolIndicators
	if err := p.Context.st.DB.Order("date asc").Find(&indicatorsLast24h,
		"\"poolAdr\" = ? AND \"date\" >= ?", foundPool.PoolAdr, time.Now().AddDate(0, 0, -1)).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "Indicators DB request error",
		})
		return
	}

	if len(indicatorsLast24h) == 0 {
		result.PriceChange24H = 0
	} else {
		totalCap24h, err := strconv.ParseFloat(indicatorsLast24h[0].TotalCap, 64)
		totalSupply24h, err := strconv.ParseFloat(indicatorsLast24h[0].TotalSupply, 64)
		if err != nil {
			response.Error(c, http.StatusBadRequest, response.E{
				Code:    response.InvalidRequest,
				Message: "ParseFloat request error",
			})
			return
		}
		if totalCap24h <= 0 || totalSupply24h <= 0 || result.CurrentPrice <= 0 {
			result.PriceChange24H = 0
		} else {
			last24HPrice := totalCap24h / totalSupply24h
			result.PriceChange24H = last24HPrice / result.CurrentPrice * 100
		}
	}
	/////

	///// Pool Transfers Data
	var investorsCount int64
	if err := p.Context.st.DB.Model(&models.PoolTransfer{}).Distinct("\"wallet\"").
		Where("\"poolAdr\" = ?", foundPool.PoolAdr).Count(&investorsCount).
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
		Where("\"poolAdr\" = ? AND \"date\" >= ?", foundPool.PoolAdr, time.Now().AddDate(0, 0, -1)).Count(&investorsLast24hCount).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "Transfers DB request error",
		})
		return
	}
	if investorsCount == 0 || investorsLast24hCount == 0 {
		result.Copiers = 0
	} else {
		result.Copiers = float64(investorsLast24hCount) / float64(investorsCount) * 100
	}

	response.Success(c, http.StatusOK, response.S{Data: result})

}
