package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	var err error
	connStr := "user=dhn password=myPassword dbname=urlshortner sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	log.Println("Connected to Postgres")

} 