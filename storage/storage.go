package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/db-go?sslmode=disable")
		if err != nil {
			panic(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Conectado a Postgres")
	})
}

func Pool() *sql.DB {
	return db
}
