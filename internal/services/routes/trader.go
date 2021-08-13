package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"time"
)

type TraderRoutes struct {
	Context *RoutesContext
}

type PoolInfoResponse struct {
	Fund              int64   `json:"fund"`
	Copiers           float64 `json:"copiers"`
	Symbol            string  `json:"symbol"`
	BasicTokenAdr     string  `json:"basicTokenAdr"`
	BasicTokenDecimal uint8   `json:"basicTokenDecimal"`
	BasicTokenSymbol  string  `json:"basicTokenSymbol"`
	CurrentPrice      string  `json:"currentPrice"`
	PriceChange24H    float64 `json:"priceChange24H"`
	TotalValueLocked  string  `json:"totalValueLocked"`
	ProfitAndLoss     string  `json:"profitAndLoss"`
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
	result.BasicTokenDecimal = foundPool.BasicTokenDecimals
	result.BasicTokenSymbol = foundPool.BasicTokenSymbol
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
	result.TotalValueLocked = indicatorLast.TotalCap

	totalCap := helpers.ToDecimal(indicatorLast.TotalCap, int(foundPool.BasicTokenDecimals))
	totalSupply := helpers.ToDecimal(indicatorLast.TotalSupply, int(foundPool.Decimals))

	if totalCap.LessThanOrEqual(decimal.NewFromInt(0)) || totalSupply.LessThanOrEqual(decimal.NewFromInt(0)) {
		result.CurrentPrice = "0"
	} else {
		result.CurrentPrice = helpers.ToWei(totalCap.Div(totalSupply), int(foundPool.BasicTokenDecimals)).String()
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
		if totalCap24h <= 0 || totalSupply24h <= 0 || result.CurrentPrice == "0" {
			result.PriceChange24H = 0
		} else {
			currentPrice, err := strconv.ParseFloat(result.CurrentPrice, 64)
			if err != nil {
				response.Error(c, http.StatusBadRequest, response.E{
					Code:    response.InvalidRequest,
					Message: "ParseFloat request error",
				})
				return
			}
			last24HPrice := totalCap24h / totalSupply24h
			result.PriceChange24H = last24HPrice / currentPrice * 100
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
