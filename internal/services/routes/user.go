package service_routes

import (
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

type UserRoutes struct {
	Context *RoutesContext
}

// @Description Update User Avatar [SIGN]
// @Summary Update User Avatar [SIGN]
// @Tags User
// @Accept  json
// @Produce  json
// @Param X-MORPH header string true "An authorization header" default(iQxX3slnRg)
// @Param wallet path string true "User wallet address"
// @Param file formData file true "Body with image"
// @Success 200 {object} response.S{}
// @Failure 400 {object} response.E
// @Router /user/{wallet}/avatar [put]
func (p *UserRoutes) PutAvatarUpdate(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("wallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}
	wallet := strings.ToLower(c.Param("wallet"))

	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "invalid FormFile",
		})
		return
	}

	path, err := filepath.Abs("public/avatars/"+wallet+".png")
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "invalid filepath",
		})
		return
	}

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "invalid Save Uploaded File",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{})

}