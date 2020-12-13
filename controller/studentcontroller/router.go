package studentcontroller

import (
	"net/http"

	std "github.com/emadghaffari/api-teacher/service/student"
	"github.com/gin-gonic/gin"
)

var (
	// Router var
	Router students = &student{}
)

// students interface
type students interface {
	Index(*gin.Context)
}

// student struct
type student struct{}

// get all courses taked by user
func (u *student) Index(c *gin.Context) {
	// create a new User
	ts, resErr := std.Service.Index()
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"error": resErr.Message()})
		return
	}
	c.JSON(http.StatusOK, ts)
}
