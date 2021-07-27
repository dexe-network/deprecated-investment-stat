package service_routes

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"strings"
)

type UserRoutes struct {
	Context *RoutesContext
}

type SignUp struct {
	Nickname string `form:"nickname" binding:"required,contains"`
	Wallet   string `form:"wallet" binding:"required,contains"`
}

// @Description User SignUp [SIGN]
// @Summary User SignUp [SIGN]
// @Tags User
// @Accept  json
// @Produce  json
// @Param x-morph header string true "An authorization header" default(iQxX3slnRg)
// @Param nickname formData string true "User nickname"
// @Param wallet formData string true "User wallet"
// @Param file formData file false "Body with image"
// @Success 200 {object} response.S{data=models.User}
// @Failure 400 {object} response.E
// @Router /user/signup [post]
func (p *UserRoutes) PostSignUp(c *gin.Context) {
	var signUp SignUp
	if err := c.ShouldBind(&signUp); err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid form data",
		})
		return
	}

	if helpers.IsValidAddress(signUp.Wallet) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}
	wallet := strings.ToLower(signUp.Wallet)

	var newUser models.User
	if err := p.Context.st.DB.First(&newUser, "\"wallet\" = ?", wallet).
		Error; (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || newUser.Id > 0 {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "User already exist",
		})
		return
	}

	file, err := c.FormFile("file")
	if err == nil {
		newUser.Avatar = wallet + ".png"
		path, err := filepath.Abs("public/avatars/" + newUser.Avatar)
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
	}

	newUser.Wallet = wallet
	newUser.Nickname = signUp.Nickname
	if err := p.Context.st.DB.Create(&newUser).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "User already exist on create",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: newUser})

}

// @Description Update User Avatar [SIGN]
// @Summary Update User Avatar [SIGN]
// @Tags User
// @Accept  json
// @Produce  json
// @Param x-morph header string true "An authorization header" default(iQxX3slnRg)
// @Param wallet path string true "User wallet address"
// @Param file formData file true "Body with image"
// @Success 200 {object} response.S{data=models.User}
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

	var newUser models.User
	if err := p.Context.st.DB.First(&newUser, "\"wallet\" = ?", wallet).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "User not exist",
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidRequest,
			Message: "invalid FormFile",
		})
		return
	}

	newUser.Avatar = wallet + ".png"
	path, err := filepath.Abs("public/avatars/" + newUser.Avatar)
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

	if err := p.Context.st.DB.Save(&newUser).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "Can't save user to DB",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: newUser})

}

// @Description Get User Info
// @Summary Get User Info
// @Tags User
// @Accept  json
// @Produce  json
// @Param wallet path string true "User wallet address"
// @Success 200 {object} response.S{data=models.User}
// @Failure 400 {object} response.E
// @Router /user/{wallet} [get]
func (p *UserRoutes) GetUserInfo(c *gin.Context) {
	if helpers.IsValidAddress(c.Param("wallet")) == false {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid wallet address",
		})
		return
	}
	wallet := strings.ToLower(c.Param("wallet"))

	var user models.User
	if err := p.Context.st.DB.First(&user, "\"wallet\" = ?", wallet).
		Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "User not exist",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: user})

}
