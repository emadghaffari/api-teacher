package redis

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	once sync.Once

	// RedisDB variable
	RedisDB iredis = &sredis{}
)

type iredis interface {
	New()
	GetDB() *redis.Client
}

type sredis struct {
	db *redis.Client
}

func (s *sredis) New() {
	once.Do(func() {
		s.db = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.Addr"),
			Password: viper.GetString("redis.Password"), // no password set
			DB:       viper.GetInt("redis.DB"),          // use default DB
		})

		fmt.Println("redis connected")
	})
}

func (s *sredis) GetDB() *redis.Client {
	return s.db
}
