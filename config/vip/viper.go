package vip

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	once sync.Once
	// Conf var
	Conf iviper = &sviper{}
)

type iviper interface {
	New()
}

// Viper struct
type sviper struct{}

// New meth for openViper
func (tt *sviper) New() {
	once.Do(func() {
		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")      // path to look for the config file in

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				log.WithFields(log.Fields{
					"error": fmt.Sprintf("Config file not found; ignore error if desired: %s", err),
				}).Fatal(fmt.Sprintf("Config file not found; ignore error if desired: %s", err))
			} else {
				// Config file was found but another error was produced
				log.WithFields(log.Fields{
					"error": fmt.Sprintf("Config file was found but another error was produced: %s", err),
				}).Fatal(fmt.Sprintf("Config file was found but another error was produced: %s", err))
			}
		}
	})
}
