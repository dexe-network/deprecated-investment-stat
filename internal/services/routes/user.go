package service_routes

import (
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRoutes struct {
	Context *RoutesContext
}

// @Description Update User Avatar
// @Summary Update User Avatar
// @Tags User
// @Accept  json
// @Produce  json
// @Param   wallet path string true "User wallet address"
// @Success 200 {object} response.S{}
// @Failure 400 {object} response.E
// @Router /user/{wallet}/avatar [post]
func (p *UserRoutes) PostAvatarUpdate(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("wallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}

	//walletAdr := strings.ToLower(c.Param("wallet"))


	response.Success(c, http.StatusOK, response.S{})

}