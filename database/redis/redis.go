package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

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
	Get(key string, dest interface{}) error
	Set(key string, value interface{}, duration time.Duration) error
	Del(key ...string) error
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

// Set meth a new key,value
func (s *redi) Set(key string, value interface{}, duration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		log.WithFields(log.Fields{
			"error": fmt.Sprintf("Marshal Error for Set New item in Redis: %s", err),
		}).Fatal(fmt.Sprintf("Marshal Error for Set New item in Redis: %s", err))
		return err
	}
	return s.db.Set(context.Background(), key, p, duration).Err()
}

// Get meth, get value with key
func (s *redi) Get(key string, dest interface{}) error {
	p, err := s.db.Get(context.Background(), key).Result()

	if p == "" {
		return fmt.Errorf("Value Not Found")
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": fmt.Sprintf("Error in Get value from Redis: %s", err),
		}).Fatal(fmt.Sprintf("Error in Get value from Redis: %s", err))
		return err
	}

	return json.Unmarshal([]byte(p), &dest)
}

func (s *redi) Del(key ...string) error {
	_, err := s.db.Del(context.Background(), key...).Result()
	if err != nil {
		return err
	}
	return nil
}
