package coursecontroller

import "github.com/gin-gonic/gin"

var (
	// Router var
	Router courses = &course{}
)

// courses interface
type courses interface {
	Store(*gin.Context)
	Index(*gin.Context)
}

// course struct
type course struct{}

// Store new course from course
func (u *course) Store(c *gin.Context) {

}

// get all courses taked by user
func (u *course) Index(c *gin.Context) {

}
