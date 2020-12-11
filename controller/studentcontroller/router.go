package studentcontroller

import (
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

}
