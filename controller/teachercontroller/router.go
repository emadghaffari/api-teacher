package teachercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/emadghaffari/api-teacher/database/redis"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/gin-gonic/gin"
)

var (
	// Router var
	Router teachers = &teacher{}
)

// teachers interface
type teachers interface {
	Store(*gin.Context)
}

// teacher struct
type teacher struct{}

// Store new course from teacher
func (u *teacher) Store(c *gin.Context) {
	resp := redis.DB.GetBy(c.Request.Header.Get("uuid"))
	r, _ := resp.Result()

	us := user.User{}
	json.Unmarshal([]byte(r), &us)
	c.JSON(http.StatusOK, map[string]string{
		"user":  us.FirstName,
		"lname": us.LastName,
		"us":    r,
	})
}
