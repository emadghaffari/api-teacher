package teachercontroller

import (
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

// get All courses offered by the teacher
func (u *teacher) Index(c *gin.Context) {
}
