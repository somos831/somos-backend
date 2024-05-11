package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dbName := os.Getenv("DB_NAME")
	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbName,
	)
	db, err := sql.Open("mysql", addr)
	if err != nil {
		connInfo := fmt.Sprintf("using connection mysql://%s\n", addr)
		log.Fatalf("connection failed: %s\n%s", err, connInfo)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to mysql database %q\n", dbName)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func Disconnect(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
