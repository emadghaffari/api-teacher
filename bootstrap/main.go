package bootstrap

import (
	"github.com/emadghaffari/api-teacher/app"
	"github.com/emadghaffari/api-teacher/config/jgr"
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

	// defer close for postgres
	defer postgres.DB.GetDB().Close()

	// defer close for redis
	defer redis.DB.GetDB().Close()

	// new jaefer nad defer close
	jgr.Tracing.New("teachers")
	defer jgr.Tracing.GetCloser().Close()

	// start application
	app.StartApplication()
}
