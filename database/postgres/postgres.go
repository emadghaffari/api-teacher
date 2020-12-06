package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	once sync.Once

	// DB variable
	DB iredis = &sredis{}
)

type iredis interface {
	New()
	GetDB() *sql.DB
}

type sredis struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func (s *sredis) New() {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=%s",
			viper.GetString("postgres.host"),
			viper.GetInt("postgres.port"),
			viper.GetString("postgres.user"),
			viper.GetString("postgres.password"),
			viper.GetString("postgres.dbname"),
			viper.GetString("postgres.sslmode"))

		var err error
		s.db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		err = s.db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("postgres connected")
	})
}

func (s *sredis) GetDB() *sql.DB {
	return s.db
}
