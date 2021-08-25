package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/nleeper/goment"
	"github.com/shopspring/decimal"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

type TraderRoutes struct {
	Context *RoutesContext
}

type PoolInfoChartData struct {
	X time.Time  `json:"x"`
	Y float64 `json:"y"`
}

type PoolInfoProfitLossByPeriod struct {
	M1  float64 `json:"m1"`
	M3  float64 `json:"m3"`
	All float64 `json:"all"`
}

type PoolInfoResponse struct {
	Fund                    int64                      `json:"fund"`
	Copiers24H              float64                    `json:"copiers24H"`
	Symbol                  string                     `json:"symbol"`
	BasicTokenAdr           string                     `json:"basicTokenAdr"`
	BasicTokenDecimal       uint8                      `json:"basicTokenDecimal"`
	BasicTokenSymbol        string                     `json:"basicTokenSymbol"`
	CurrentPrice            string                     `json:"currentPrice"`
	PriceChange24H          float64                    `json:"priceChange24H"`
	TotalValueLocked        string                     `json:"totalValueLocked"`
	ProfitAndLoss           float64                    `json:"profitAndLoss"`
	PersonalFundsLocked     string                     `json:"personalFundsLocked"`
	InvestorsFundsLocked    string                     `json:"investorsFundsLocked"`
	PersonalFundsLocked24H  float64                    `json:"personalFundsLocked24H"`
	InvestorsFundsLocked24H float64                    `json:"investorsFundsLocked24H"`
	ProfitAndLossByPeriod   PoolInfoProfitLossByPeriod `json:"profitAndLossByPeriod"`
	ProfitAndLossChart      []PoolInfoChartData        `json:"profitAndLossChart"`
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
		response.Error(
			c, http.StatusBadRequest, response.E{
				Code:    response.InvalidJSONBody,
				Message: "invalid pool Address",
			},
		)
		return
	}

	////// Pool Data
	poolAddress := c.Param("poolAddress")
	var foundPool models.Pool
	if err := p.Context.st.DB.First(&foundPool, "LOWER(\"poolAdr\") = LOWER(?)", poolAddress).
		Error; err != nil {
		response.Error(
			c, http.StatusBadRequest, response.E{
				Code:    response.InvalidRequest,
				Message: "Pool not found",
			},
		)
		return
	}
	//////

	////// Indicators Data
	var indicatorLast models.PoolIndicators
	if err := p.Context.st.DB.Order("date desc").First(
		&indicatorLast,
		"\"poolAdr\" = ?", foundPool.PoolAdr,
	).
		Error; err != nil {
		response.Error(
			c, http.StatusBadRequest, response.E{
				Code:    response.InvalidRequest,
				Message: "Indicators DB request error",
			},
		)
		return
	}

	investorsFundsLocked,
	personalFundsLocked,
	totalValueLocked,
	currentPrice,
	profitAndLoss := getPoolInfoIndicatorData(&indicatorLast, &foundPool)

	////// Indicators Last 24 Data
	var indicatorsLast24h []models.PoolIndicators
	if err := p.Context.st.DB.Order("date asc").Find(
		&indicatorsLast24h,
		"\"poolAdr\" = ? AND \"date\" >= ?", foundPool.PoolAdr, time.Now().AddDate(0, 0, -1),
	).
		Error; err != nil {
		response.Error(
			c, http.StatusBadRequest, response.E{
				Code:    response.InvalidRequest,
				Message: "Indicators DB request error",
			},
		)
		return
	}
	priceChange24H, personalFundsLocked24H, investorsFundsLocked24H := getPoolInfoIndicatorLast24Data(
		indicatorsLast24h,
		currentPrice,
		investorsFundsLocked,
		personalFundsLocked,
		c,
	)
	/////

	///// Pool Transfers Data
	var investorsCount int64
	if err := p.Context.st.DB.Model(&models.PoolTransfer{}).Distinct("\"wallet\"").
		Where("\"poolAdr\" = ?", foundPool.PoolAdr).Count(&investorsCount).
		Error; err != nil {
		response.Error(
			c, http.StatusBadRequest, response.E{
				Code:    response.InvalidRequest,
				Message: "Transfers DB request error",
			},
		)
		return
	}
	fund := investorsCount

	var investorsLast24hCount int64
	if err := p.Context.st.DB.Model(&models.PoolTransfer{}).Distinct("\"wallet\"").
		Where(
			"\"poolAdr\" = ? AND \"date\" >= ?",
			foundPool.PoolAdr,
			time.Now().AddDate(0, 0, -1),
		).Count(&investorsLast24hCount).
		Error; err != nil {
		response.Error(
			c, http.StatusBadRequest, response.E{
				Code:    response.InvalidRequest,
				Message: "Transfers DB request error",
			},
		)
		return
	}
	copiers := getPoolInfoInvestorsLast24hCount(investorsCount, investorsLast24hCount)

	var indicatorsAll []models.PoolIndicators
	if err := p.Context.st.DB.Order("date asc").Find(
		&indicatorsAll,
		"\"poolAdr\" = ?", foundPool.PoolAdr,
	).
		Error; err != nil {
		response.Error(
			c, http.StatusBadRequest, response.E{
				Code:    response.InvalidRequest,
				Message: "Indicators DB request error",
			},
		)
		return
	}
	profitAndLossByPeriod := getProfitAndLossByPeriod(indicatorsAll, &foundPool)
	profitAndLossChart := getProfitAndLossChart(indicatorsAll)

	result := &PoolInfoResponse{
		Fund:                    fund,
		Copiers24H:              copiers,
		Symbol:                  foundPool.Symbol,
		BasicTokenAdr:           foundPool.BasicTokenAdr,
		BasicTokenDecimal:       foundPool.BasicTokenDecimals,
		BasicTokenSymbol:        foundPool.BasicTokenSymbol,
		CurrentPrice:            currentPrice,
		PriceChange24H:          priceChange24H,
		TotalValueLocked:        totalValueLocked,
		ProfitAndLoss:           profitAndLoss,
		PersonalFundsLocked:     personalFundsLocked,
		InvestorsFundsLocked:    investorsFundsLocked,
		PersonalFundsLocked24H:  personalFundsLocked24H,
		InvestorsFundsLocked24H: investorsFundsLocked24H,
		ProfitAndLossByPeriod:   profitAndLossByPeriod,
		ProfitAndLossChart:      profitAndLossChart,
	}

	response.Success(c, http.StatusOK, response.S{Data: result})

}

func getProfitAndLossChart(indicatorsAll []models.PoolIndicators) (poolInfoChartData []PoolInfoChartData) {
	if len(indicatorsAll) == 0 {
		return
	}

	for _, indicators := range indicatorsAll {
		price, _ := decimal.NewFromString(indicators.PoolTokenPrice)
		if price.GreaterThan(decimal.NewFromInt(0)) {
			profitAndLoss, _ := price.Mul(decimal.NewFromInt(100)).
				Div(decimal.NewFromInt(1)).
				Sub(decimal.NewFromInt(100)).Float64()
			poolInfoChartData = append(
				poolInfoChartData,
				PoolInfoChartData{X: indicators.Date.UTC(), Y: profitAndLoss},
			)
		}
	}
	return
}

func getProfitAndLossByPeriod(
	indicatorsAll []models.PoolIndicators,
	foundPool *models.Pool,
) (profitAndLossByPeriod PoolInfoProfitLossByPeriod) {
	if len(indicatorsAll) == 0 {
		return
	}

	currentTime, _ := goment.New()
	m1Time := currentTime.Subtract(1, "month").ToUnix()
	m3Time := currentTime.Subtract(3, "month").ToUnix()

	var m1 []models.PoolIndicators
	var m3 []models.PoolIndicators

	for _, indicator := range indicatorsAll {
		indicatorDate := indicator.Date.Unix()
		if indicatorDate > m1Time {
			m1 = append(m1, indicator)
			continue
		}

		if indicatorDate > m3Time {
			m3 = append(m3, indicator)
			continue
		}
	}

	// Latest
	latestIndicator := indicatorsAll[len(indicatorsAll)-1]
	latestTotalCap := helpers.ToDecimal(latestIndicator.TotalCap, int(foundPool.BasicTokenDecimals))
	latestTotalSupply := helpers.ToDecimal(latestIndicator.TotalSupply, int(foundPool.Decimals))
	latestPrice := latestTotalCap.Div(latestTotalSupply)
	profitAndLossByPeriod.All, _ = latestPrice.Mul(decimal.NewFromInt(100)).
		Div(decimal.NewFromInt(1)).
		Sub(decimal.NewFromInt(100)).Float64()

	// m1
	if len(m1) > 0 {
		m1Indicator := m1[0]
		m1totalCap := helpers.ToDecimal(m1Indicator.TotalCap, int(foundPool.BasicTokenDecimals))
		m1totalSupply := helpers.ToDecimal(m1Indicator.TotalSupply, int(foundPool.Decimals))
		m1Price := m1totalCap.Div(m1totalSupply)
		profitAndLossByPeriod.M1, _ = latestPrice.Mul(decimal.NewFromInt(100)).
			Div(m1Price).
			Sub(decimal.NewFromInt(100)).Float64()
	} else {
		profitAndLossByPeriod.M1 = 0
	}

	// m3
	if len(m3) > 0 {
		m3Indicator := m3[0]
		m3totalCap := helpers.ToDecimal(m3Indicator.TotalCap, int(foundPool.BasicTokenDecimals))
		m3totalSupply := helpers.ToDecimal(m3Indicator.TotalSupply, int(foundPool.Decimals))
		m3Price := m3totalCap.Div(m3totalSupply)
		profitAndLossByPeriod.M3, _ = latestPrice.Mul(decimal.NewFromInt(100)).
			Div(m3Price).
			Sub(decimal.NewFromInt(100)).Float64()
	} else {
		profitAndLossByPeriod.M3 = 0
	}

	return
}

func getPoolInfoInvestorsLast24hCount(investorsCount int64, investorsLast24hCount int64) (copiers float64) {
	if investorsCount == 0 || investorsLast24hCount == 0 {
		copiers = 0
	} else {
		copiers = float64(investorsLast24hCount) / float64(investorsCount) * 100
	}
	return
}

func getPoolInfoIndicatorLast24Data(
	indicatorsLast24h []models.PoolIndicators,
	currentPrice string,
	investorsFundsLocked string,
	personalFundsLocked string,
	c *gin.Context,
) (priceChange24H float64, personalFundsLocked24H float64, investorsFundsLocked24H float64) {
	if len(indicatorsLast24h) == 0 {
		priceChange24H = 0
		personalFundsLocked24H = 0
		investorsFundsLocked24H = 0
	} else {
		indicatorData := indicatorsLast24h[0]
		totalCap24h, err := strconv.ParseFloat(indicatorData.TotalCap, 64)
		totalSupply24h, err := strconv.ParseFloat(indicatorData.TotalSupply, 64)
		latestInvestorsFundsLocked, err := strconv.ParseFloat(investorsFundsLocked, 64)
		latestPersonalFundsLocked, err := strconv.ParseFloat(personalFundsLocked, 64)
		if err != nil {
			response.Error(
				c, http.StatusBadRequest, response.E{
					Code:    response.InvalidRequest,
					Message: "ParseFloat request error",
				},
			)
			return
		}
		if totalCap24h <= 0 || totalSupply24h <= 0 || currentPrice == "0" {
			priceChange24H = 0
		} else {
			currentPrice, err := strconv.ParseFloat(currentPrice, 64)
			if err != nil {
				response.Error(
					c, http.StatusBadRequest, response.E{
						Code:    response.InvalidRequest,
						Message: "ParseFloat request error",
					},
				)
				return
			}
			last24HPrice := totalCap24h / totalSupply24h
			priceChange24H = last24HPrice / currentPrice * 100
		}

		//// Calculate personalFundsLocked24H and investorsFundsLocked24H
		// Parse String to Int
		oldTotalCapInt := new(big.Int)
		oldTotalCapInt.SetString(indicatorData.TotalCap, 10)
		oldTraderBasicTokensDeposited := new(big.Int)
		oldTraderBasicTokensDeposited.SetString(indicatorData.TraderBasicTokensDeposited, 10)
		//

		oldestInvestorsFundsLocked := float64(
			big.NewInt(0).Sub(
				oldTotalCapInt,
				oldTraderBasicTokensDeposited,
			).Int64(),
		)
		oldestPersonalFundsLocked := float64(oldTraderBasicTokensDeposited.Int64())

		if oldestInvestorsFundsLocked <= 0 {
			investorsFundsLocked24H = 0
		} else {
			investorsFundsLocked24H = oldestInvestorsFundsLocked / latestInvestorsFundsLocked * 100
		}

		if oldestPersonalFundsLocked <= 0 {
			personalFundsLocked24H = 0
		} else {
			personalFundsLocked24H = oldestPersonalFundsLocked / latestPersonalFundsLocked * 100
		}
	}
	return
}

func getPoolInfoIndicatorData(indicatorLast *models.PoolIndicators, foundPool *models.Pool) (
	investorsFundsLocked string,
	personalFundsLocked string,
	totalValueLocked string,
	currentPrice string,
	profitAndLoss float64,
) {
	// Parse String to Int
	totalCapInt := new(big.Int)
	totalCapInt.SetString(indicatorLast.TotalCap, 10)
	traderBasicTokensDeposited := new(big.Int)
	traderBasicTokensDeposited.SetString(indicatorLast.TraderBasicTokensDeposited, 10)
	//

	investorsFundsLocked = big.NewInt(0).Sub(
		totalCapInt,
		traderBasicTokensDeposited,
	).String()
	personalFundsLocked = traderBasicTokensDeposited.String()
	totalValueLocked = indicatorLast.TotalCap

	totalCap := helpers.ToDecimal(indicatorLast.TotalCap, int(foundPool.BasicTokenDecimals))
	totalSupply := helpers.ToDecimal(indicatorLast.TotalSupply, int(foundPool.Decimals))

	if totalCap.LessThanOrEqual(decimal.NewFromInt(0)) || totalSupply.LessThanOrEqual(decimal.NewFromInt(0)) {
		currentPrice = "0"
		// need improve/investigate
		profitAndLoss = float64(-100)
	} else {
		currentPriceRaw := totalCap.Div(totalSupply)
		currentPrice = helpers.ToWei(currentPriceRaw, int(foundPool.BasicTokenDecimals)).String()
		// PL will be correct when start token price 1 token = 1 baseToken
		profitAndLoss, _ = currentPriceRaw.Mul(decimal.NewFromInt(100)).
			Div(decimal.NewFromInt(1)).
			Sub(decimal.NewFromInt(100)).Float64()
	}
	return
}
