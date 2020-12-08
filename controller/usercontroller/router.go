package usercontroller

import (
	"net/http"

	model "github.com/emadghaffari/api-teacher/model/user"
	service "github.com/emadghaffari/api-teacher/service/user"
	cryptoutils "github.com/emadghaffari/api-teacher/utils/cryptoUtils"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
)

var (
	// Router var
	Router users = &user{}
)

// user interface
type users interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

// user struct
type user struct{}

func (u *user) Login(c *gin.Context) {
	us := model.User{}

	// Bind the request.Body to user
	if err := c.ShouldBindJSON(&us); err != nil {
		resErr := errors.HandlerBadRequest("Invalid JSON Body.")
		c.JSON(resErr.Status(), resErr.Message())
		return
	}
	us.Password = cryptoutils.GetMD5(us.Password)

	// Login
	ts, resErr := service.Service.Login(us)
	if resErr != nil {
		c.JSON(resErr.Status(), resErr)
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func (u *user) Register(c *gin.Context) {
	us := model.User{}

	// Bind the request.Body to user
	if err := c.ShouldBindJSON(&us); err != nil {
		resErr := errors.HandlerBadRequest("Invalid JSON Body.")
		c.JSON(resErr.Status(), resErr.Message())
		return
	}
	us.Password = cryptoutils.GetMD5(us.Password)

	// create a new User
	ts, resErr := service.Service.Register(us)
	if resErr != nil {
		c.JSON(resErr.Status(), resErr)
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}
