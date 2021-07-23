package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type NonceRoutes struct {
	Context *RoutesContext
}

// @Description Get Sign Nonce
// @Summary Get Sign Nonce
// @Tags Nonce
// @Accept  json
// @Produce  json
// @Param   wallet path string true "User wallet address"
// @Success 200 {object} response.S{data=GetNonceResponse}
// @Failure 400 {object} response.E
// @Router /nonce/{wallet} [get]
func (p *NonceRoutes) GetNonce(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("wallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}

	walletAdr := strings.ToLower(c.Param("wallet"))

	var nonce models.Nonce
	if err := p.Context.st.DB.FirstOrCreate(&nonce,
		&models.Nonce{
			Wallet: walletAdr,
		}).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	randomNonce := helpers.RandomNonce()
	nonce.Nonce = randomNonce

	p.Context.st.DB.Save(&nonce)

	response.Success(c, http.StatusOK, response.S{Data: GetNonceResponse{Nonce: randomNonce}})

}

type GetNonceResponse struct {
	Nonce int `json:"nonce"`
}
