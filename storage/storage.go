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
		db, err = sql.Open("postgres", "postgres://postgres:admin@localhost:5432/db-go?sslmode=disable")
		//db, err = sql.Open("postgres", "postgres://postgres:xxx@localhost:5432/db-go?sslmode=disable")

		if err != nil {
			log.Fatalf("antes de panico: %v", err)
			//panic(err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Conectado a Postgres")
	})
}

func Pool() *sql.DB {
	return db
}
