package config

import (
	"database/sql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/newsapp"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("database tidak terkoneksi")
	}

	return db, nil
}
