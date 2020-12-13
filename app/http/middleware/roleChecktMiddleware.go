package middleware

import (
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
	Check(role string) gin.HandlerFunc
}

// role struct
type role struct{}

func (ac *role) Check(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("role", role)
		ac.check(c)
	}
}

// check role
func (ac *role) check(c *gin.Context) {
	// get uuid from header
	data := c.Request.Header.Get("uuid")
	if data == "" {
		Middleware.RespondWithErrorJSON(c, http.StatusBadRequest, "invalid uuid for user.")
		return
	}

	// get data from redis and unmarshal data
	// new user struct
	us := user.User{}
	if err := redis.DB.Get(data, &us); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, ext := c.Get("role")
	if !ext {
		Middleware.RespondWithErrorJSON(c, http.StatusNotAcceptable, fmt.Sprintf("user not allowed to Access."))
		return
	}

	// if user role != seted role
	if role != us.Role.Name {
		Middleware.RespondWithErrorJSON(c, http.StatusNotAcceptable, fmt.Sprintf("user not allowed to Access"))
		return
	}

	// set struct and go to next!
	user.Model.Set(&us)
	c.Next()
}
