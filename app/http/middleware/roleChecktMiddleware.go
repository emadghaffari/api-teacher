package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emadghaffari/api-teacher/database/redis"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/gin-gonic/gin"
)

var (
	// Role handler
	Role rolMiddleware = &role{}
)

// rolMiddleware interface
type rolMiddleware interface {
	Check(c *gin.Context)
	SetRole(role string)
}

// role struct
type role struct {
	Role string
}

// check role
func (ac *role) Check(c *gin.Context) {

	// get uuid from header
	data := c.Request.Header.Get("uuid")
	if data == "" {
		Middleware.RespondWithErrorJSON(c, http.StatusBadRequest, "invalid uuid for user.")
	}

	// get data from redis and unmarshal data
	r, err := redis.DB.GetBy(data).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// new user struct
	us := user.User{}
	if err := json.Unmarshal([]byte(r), &us); err != nil {
		Middleware.RespondWithErrorJSON(c, http.StatusBadRequest, fmt.Sprintf("error in unmarshal user: %v", err))
	}

	// if user role != seted role
	if ac.Role != us.Role.Name {
		Middleware.RespondWithErrorJSON(c, http.StatusNotAcceptable, fmt.Sprintf("user not allowed to Access"))
	}

	// set struct and go to next!
	user.Model.Set(&us)
	c.Next()
}

// SetRole for check by middleware
func (ac *role) SetRole(role string) {
	ac.Role = role
}
