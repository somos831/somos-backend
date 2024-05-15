package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

var ErrEventNotFound = errors.New("event not found")

type Event struct {
	Id              int     `json:"id"`
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	OrganizationId  *int    `json:"organization_id"`
	ImageId         *int    `json:"image_id"`
	LocationId      *int    `json:"location_id"`
	LocationDetails *string `json:"location_details"`
	Price           float32 `json:"price"`
	CategoryId      int     `json:"category_id"`
	AdditionalInfo  *string `json:"additional_info"`
	AdditionalUrl   *string `json:"additional_url"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// FindEventById finds an event in db by its id eventId.
func FindEventById(ctx context.Context, db *sql.DB, eventId int) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.QueryRowContext(ctx, query, eventId)

	var event Event
	err := row.Scan(
		&event.Id,
		&event.Title,
		&event.Description,
		&event.OrganizationId,
		&event.ImageId,
		&event.LocationId,
		&event.LocationDetails,
		&event.Price,
		&event.CategoryId,
		&event.AdditionalInfo,
		&event.AdditionalUrl,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEventNotFound
		}
		log.Printf("failed to find event by id: %s\nid: %d\n", err, eventId)

		return nil, err
	}

	return &event, nil
}

// InsertEvent inserts event into db. The id of the event inserted is returned.
func InsertEvent(ctx context.Context, db *sql.DB, event Event) (int, error) {
	query := `
		INSERT INTO events (
			title,
			description,
			organization_id,
			image_id,
			location_id,
			location_details,
			price,
			category_id,
			additional_info,
			additional_url
		) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
	`
	result, err := db.ExecContext(ctx, query,
		event.Title,
		event.Description,
		event.OrganizationId,
		event.ImageId,
		event.LocationId,
		event.LocationDetails,
		event.Price,
		event.CategoryId,
		event.AdditionalInfo,
		event.AdditionalUrl,
	)

	if err != nil {
		log.Printf("failed to insert event: %s\n", err)
		return 0, err
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		log.Printf("failed to retreive event id: %s\n", err)
		return 0, err
	}

	return int(eventId), err
}

// UpdateEvent updates an event in db.
func UpdateEvent(ctx context.Context, db *sql.DB, event Event) (*Event, error) {
	if _, err := FindEventById(ctx, db, event.Id); err != nil {
		return nil, err
	}

	query := `
		UPDATE events SET
			title = ?,
			description = ?,
			organization_id = ?,
			image_id = ?,
			location_id = ?,
			location_details = ?,
			price = ?,
			category_id = ?,
			additional_info = ?,
			additional_url = ?
		WHERE id = ?
	`
	_, err := db.ExecContext(ctx, query,
		event.Title,
		event.Description,
		event.OrganizationId,
		event.ImageId,
		event.LocationId,
		event.LocationDetails,
		event.Price,
		event.CategoryId,
		event.AdditionalInfo,
		event.AdditionalUrl,
		event.Id,
	)

	if err != nil {
		log.Printf("failed to update event: %s\n", err)
		return nil, err
	}

	query = `SELECT created_at, updated_at FROM events WHERE id = ?`
	row := db.QueryRowContext(ctx, query, event.Id)
	if err := row.Scan(&event.CreatedAt, &event.UpdatedAt); err != nil {
		return nil, err
	}

	return &event, nil
}

// DeleteEvent deletes an event using eventId.
func DeleteEvent(ctx context.Context, db *sql.DB, eventId int) error {
	if _, err := FindEventById(ctx, db, eventId); err != nil {
		return err
	}

	query := `DELETE FROM events WHERE id = ?`
	_, err := db.ExecContext(ctx, query, eventId)
	if err != nil {
		log.Printf("failed to delete event: %s\nid: %d\n", err, eventId)
		return err
	}

	return nil
}
