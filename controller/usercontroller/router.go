package usercontroller

import (
	"github.com/gin-gonic/gin"
)

var (
	// Router var
	Router users = &user{}
)

// user interface
type users interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

// user struct
type user struct{}

func (u *user) Login(c *gin.Context) {

}

func (u *user) Register(c *gin.Context) {

}
