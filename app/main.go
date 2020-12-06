package app

import (
	"github.com/emadghaffari/api-teacher/app/http/router"
	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/database/redis"
)

// StartApplication func
// with this func we can start application
func StartApplication() {
	defer postgres.DB.GetDB().Close()
	defer redis.DB.GetDB().Close()
	router.Map()
}
