package coursecontroller

import (
	"net/http"

	service "github.com/emadghaffari/api-teacher/service/course"
	"github.com/gin-gonic/gin"
)

var (
	// Router var
	Router courses = &course{}
)

// courses interface
type courses interface {
	Index(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Take(c *gin.Context)
}

// course struct
type course struct{}

// get all courses
func (u *course) Index(c *gin.Context) {
	// create a new User
	ts, resErr := service.Service.Index()
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"error": resErr.Message()})
		return
	}
	c.JSON(http.StatusOK, ts)
}

// Store new course
func (u *course) Store(c *gin.Context) {
}

// Store new course
func (u *course) Update(c *gin.Context) {

}

// take a course
func (u *course) Take(c *gin.Context) {

}
