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
	Store(*gin.Context)
}

// teacher struct
type teacher struct{}

// Store new course from teacher
func (u *teacher) Store(c *gin.Context) {
}
