package app

import (
	"github.com/emadghaffari/api-teacher/app/http/router"
)

// StartApplication func
// with this func we can start application
func StartApplication() {
	router.Map()
}
