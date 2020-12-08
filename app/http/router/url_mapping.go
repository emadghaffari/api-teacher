package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// Router for gin
	router = gin.Default()
)

func init() {
	gin.SetMode(viper.GetString("gin.mode"))
}

// Map all routes
func Map() {
	user()
	course()
	student()
	teacher()

	// map all urls
	router.Run(viper.GetString("gin.port"))
}
