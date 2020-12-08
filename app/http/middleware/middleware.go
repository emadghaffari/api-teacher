package middleware

import (
	"github.com/gin-gonic/gin"
)

var(
	// Middleware handler
	Middleware middlewareInterface = &middleware{}
)


type middlewareInterface interface {
	RespondWithErrorJSON(c *gin.Context, code int, message interface{})
	RespondWithError(c *gin.Context, code int, err error)
}

type middleware struct{}

func (mi *middleware) RespondWithErrorJSON(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func (mi *middleware) RespondWithError(c *gin.Context, code int, err error) {
	c.AbortWithError(code,err)
}