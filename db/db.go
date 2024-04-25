package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(user, password, name string) *sql.DB {
	connInfo := fmt.Sprintf("%s:%s@/%s", user, password, name)
	db, err := sql.Open("mysql", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("connected to mysql database %q\n", name)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func Disconnect(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func ExecuteSchemaFromFile(db *sql.DB, filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read SQL schema file: %s\n", err)
	}

	statements := strings.Split(string(content), ";")
	for _, stmt := range statements {
		trimmed := strings.TrimSpace(stmt)
		if trimmed == "" {
			continue
		}

		_, err := db.Exec(trimmed)
		if err != nil {
			log.Fatalf("failed to execute SQL statement: %s\n%s\n", err, trimmed)
		}
	}

	log.Println("SQL schema executed successfully")
}
