package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	user     string
	password string
	host     string
	port     int
	name     string
}

func DefaultConfig(name string) Config {
	return Config{
		user:     "root",
		password: "root",
		host:     "127.0.0.1",
		port:     3306,
		name:     name,
	}
}

func (cfg Config) SetUser(val string) Config {
	cfg.user = val
	return cfg
}

func (cfg Config) SetPassword(val string) Config {
	cfg.password = val
	return cfg
}

func (cfg Config) SetHost(val string) Config {
	cfg.host = val
	return cfg
}

func (cfg Config) SetPort(val int) Config {
	cfg.port = val
	return cfg
}

func (cfg Config) SetName(val string) Config {
	cfg.name = val
	return cfg
}

func (cfg Config) dataSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.user,
		cfg.password,
		cfg.host,
		cfg.port,
		cfg.name,
	)
}

func Connect(cfg Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.dataSource())
	if err != nil {
		log.Fatalf("connection failed: %s\nconfig: mysql://%s\n", err, cfg.dataSource())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to mysql database %q\n", cfg.name)

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
