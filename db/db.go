package db

import (
	"database/sql"
	"fmt"
	"log"
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

func CreateTables(db *sql.DB) {
	_, err := db.Query(`
		CREATE TABLE IF NOT EXISTS categories (
 			id INT NOT NULL AUTO_INCREMENT,
 			name VARCHAR(50) NOT NULL,
      		PRIMARY KEY (id)
        );
    `)
	if err != nil {
		log.Fatalf("failed to create table 'categories': %s\n", err)
	}
	log.Println("table 'categories' found or created")
	_, err = db.Query(`
		CREATE TABLE IF NOT EXISTS events (
    		id INT NOT NULL AUTO_INCREMENT,
     		name VARCHAR(50) NOT NULL,
      		description VARCHAR(1000),
       		category_id INT NOT NULL,
        	location VARCHAR(200),
        	PRIMARY KEY (id),
        	FOREIGN KEY (category_id) REFERENCES categories(id)
    	);
    `)
	if err != nil {
		log.Fatalf("failed to create table 'events': %s\n", err)
	}
	log.Println("table 'events' found or created")
}
