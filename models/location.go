package models

import (
	"context"
	"database/sql"
	"log"
)

type Location struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	URL     string `json:"url"`
}

func InsertLocation(ctx context.Context, db *sql.DB, loc Location) (int, error) {
	query := "INSERT INTO locations (name, address, url) VALUES (?, ?, ?)"

	result, err := db.ExecContext(ctx, query, loc.Name, loc.Address, loc.URL)
	if err != nil {
		log.Printf("failed to insert location details: %s\n", err)
		return 0, err
	}

	locationID, err := result.LastInsertId()
	if err != nil {
		log.Printf("failed to retreive location id: %s\n", err)
		return 0, err
	}

	return int(locationID), err
}
