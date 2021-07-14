package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type TradesRoutes struct {
	Context *RoutesContext
}

type responseTrade struct {
	Id           uint      `json:"id"`
	TraderPool   string    `json:"traderPool"`
	AmountIn     string    `json:"amountIn"`
	AmountOutMin string    `json:"amountOutMin"`
	Path         []string  `json:"path"`
	Deadline     string    `json:"deadline"`
	Date         time.Time `json:"date"`
	BlockNumber  int64     `json:"blockNumber"`
	Tx           string    `json:"tx"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// @Description Get Trades By PoolAdr
// @Summary Get Trades By PoolAdr
// @Tags Trades
// @Accept  json
// @Produce  json
// @Param   traderPool path string true "traderPool Address"
// @Success 200 {object} response.S{data=[]responseTrade}
// @Failure 400 {object} response.E
// @Router /trades/{traderPool} [get]
func (p *TradesRoutes) GetTradesByPoolAdr(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("traderPool")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}

	var pools []models.Trade
	if err := p.Context.st.DB.Find(&pools, "\"traderPool\" = ?", c.Param("traderPool")).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	var result []responseTrade
	for _, pool := range pools {
		pathRawValue, _ := pool.Path.Value()
		parsedPathString := strings.TrimSuffix(strings.TrimPrefix(pathRawValue.(string), "{"), "}")
		path := strings.Split(parsedPathString, ",")

		item := &responseTrade{
			pool.Id,
			pool.TraderPool,
			pool.AmountIn,
			pool.AmountOutMin,
			path,
			pool.Deadline,
			pool.Date,
			pool.BlockNumber,
			pool.Tx,
			pool.CreatedAt,
			pool.UpdatedAt,
		}

		result = append(result, *item)
	}

	response.Success(c, http.StatusOK, response.S{Data: result})

}
