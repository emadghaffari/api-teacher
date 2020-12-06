package log

import (
	"fmt"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	once sync.Once
	// Conf variable instance of LogrusInterface
	Conf logrusInterface = &logrusStruct{}
)

type logrusInterface interface {
	New()
}
type logrusStruct struct{}

// Configs func
func (l logrusStruct) New() {
	once.Do(func() {
		if viper.GetString("Environment") == "production" {
			// Log as JSON instead of the default ASCII formatter.
			f, _ := os.OpenFile(fmt.Sprintf("logs/%s.log", time.Now().Local().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
			log.SetFormatter(&log.JSONFormatter{})
			log.SetOutput(f)

		} else {
			// The TextFormatter is default, you don't actually have to do this.
			log.SetFormatter(&log.TextFormatter{})

			// Output to stdout instead of the default stderr
			log.SetOutput(os.Stdout)

		}

		// Only log the warning severity or above.
		log.SetLevel(log.InfoLevel)
	})

}
