package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoRoutes struct {
	Context *RoutesContext
}

// @Description Get Global Token Whitelist
// @Summary Get Global Token Whitelist
// @Tags Info
// @Accept  json
// @Produce  json
// @Success 200 {object} response.S{data=[]models.GlobalTokenWhitelist}
// @Failure 400 {object} response.E
// @Router /info/global-token-whitelist [get]
func (p *InfoRoutes) GetGlobalTokenWhitelist(c *gin.Context) {
	var tokens []models.GlobalTokenWhitelist
	if err := p.Context.st.DB.Find(&tokens).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: tokens})

}
