package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/emadghaffari/api-teacher/config/token"
	"github.com/gin-gonic/gin"
)

var (
	// AccessToken handler
	AccessToken actMiddleware = &act{}
)

// actMiddleware interface
type actMiddleware interface {
	CheckMiddleware(c *gin.Context)
}

type act struct{}

func (ac *act) CheckMiddleware(c *gin.Context) {
	bearToken := c.Request.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		Middleware.RespondWithErrorJSON(c, http.StatusBadRequest, "the access token is invalid.")
	}
	resp, err := token.Conf.VerifyToken(strArr[1])
	if err != nil {
		Middleware.RespondWithErrorJSON(c, http.StatusBadRequest, fmt.Sprintf("the access token is invalid. %v", err))
	} else {
		c.Request.Header.Set("userID", fmt.Sprint(resp.UserID))
		c.Request.Header.Set("accessUUID", string(resp.AccessUUID))
		c.Next()
	}
}
