package teachercontroller

import (
	"net/http"

	tech "github.com/emadghaffari/api-teacher/service/teacher"
	"github.com/gin-gonic/gin"
)

var (
	// Router var
	Router teachers = &teacher{}
)

// teachers interface
type teachers interface {
	Index(*gin.Context)
}

// teacher struct
type teacher struct{}

// get all courses taked by user
func (u *teacher) Index(c *gin.Context) {
	// create a new User
	ts, resErr := tech.Service.Index()
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"error": resErr.Message()})
		return
	}
	c.JSON(http.StatusOK, ts)
}
