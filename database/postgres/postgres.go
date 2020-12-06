package postgres

import (
	"database/sql"
	"fmt"
	"sync"
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
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		var err error
		s.db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		fmt.Println("postgres connected")
	})
}

func (s *sredis) GetDB() *sql.DB {
	return s.db
}
