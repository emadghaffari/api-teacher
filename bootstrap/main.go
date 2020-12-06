package bootstrap

import (
	"github.com/emadghaffari/api-teacher/app"
	"github.com/emadghaffari/api-teacher/config/log"
	"github.com/emadghaffari/api-teacher/config/vip"
)

// manage configs
func init() {
	vip.Conf.New()
	log.Conf.New()
}

// Boot func call start application
func Boot() {
	app.StartApplication()
}
