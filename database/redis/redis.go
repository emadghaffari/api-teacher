package redis

import (
	"context"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	once sync.Once

	// DB variable
	DB rs = &redi{}
)

type rs interface {
	New()
	GetDB() *redis.Client
}

type redi struct {
	db *redis.Client
}

func (s *redi) New() {
	once.Do(func() {
		s.db = redis.NewClient(&redis.Options{
			Addr: viper.GetString("redis.Addr"),
			// Password: viper.GetString("redis.Password"), // no password set
			DB: viper.GetInt("redis.DB"), // use default DB
		})

		if err := s.db.Ping(context.Background()).Err(); err != nil {
			log.WithFields(log.Fields{
				"error": fmt.Sprintf("Config redis database Error: %s", err),
			}).Fatal(fmt.Sprintf("Config redis database Error: %s", err))
		}

		fmt.Println("redis connected")
	})
}

func (s *redi) GetDB() *redis.Client {
	return s.db
}
