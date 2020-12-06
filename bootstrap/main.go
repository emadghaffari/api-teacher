package bootstrap

import (
	"github.com/emadghaffari/api-teacher/app"
	"github.com/emadghaffari/api-teacher/config/log"
	"github.com/emadghaffari/api-teacher/config/vip"
	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/database/redis"
)

// manage configs
func init() {
	vip.Conf.New()
	log.Conf.New()
	postgres.DB.New()
	redis.DB.New()
}

// Boot func call start application
func Boot() {
	app.StartApplication()
}
