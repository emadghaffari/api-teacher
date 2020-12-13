package helper

import (
	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
)

// Bind func for bind gin controllers
func Bind(c *gin.Context, cs interface{}) errors.ResError {
	// Bind the request.Body to user
	if err := c.ShouldBindJSON(&cs); err != nil {
		return errors.HandlerBadRequest("Invalid JSON Body.")

	}
	return nil
}
