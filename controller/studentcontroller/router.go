package studentcontroller

import "github.com/gin-gonic/gin"

var (
	// Router var
	Router students = &student{}
)

// students interface
type students interface {
	Store(*gin.Context)
	Index(*gin.Context)
}

// student struct
type student struct{}

// Store new course from student
func (u *student) Store(c *gin.Context) {

}

// get all courses taked by user
func (u *student) Index(c *gin.Context) {

}
