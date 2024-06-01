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
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"end_date"`
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
	IsVisible       bool    `json:"is_visible"`
	ContactInfo     string  `json:"contact_info"`
}

// FindNRecentEvents finds the n most recent events in db and orders them from
// most recent to least recent.
func FindNRecentEvents(ctx context.Context, db *sql.DB, n int) ([]Event, error) {
	query := `SELECT * FROM events ORDER BY date DESC LIMIT ?`
	rows, err := db.QueryContext(ctx, query, n)
	if err != nil {
		return nil, err
	}

	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.Id,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
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
			&event.IsVisible,
			&event.ContactInfo,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
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
		&event.StartDate,
		&event.EndDate,
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
		&event.IsVisible,
		&event.ContactInfo,
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
			start_date,
			end_date,
			organization_id,
			image_id,
			location_id,
			location_details,
			price,
			category_id,
			additional_info,
			additional_url,
			contact_info
		) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
	`
	result, err := db.ExecContext(ctx, query,
		event.Title,
		event.Description,
		event.StartDate,
		event.EndDate,
		event.OrganizationId,
		event.ImageId,
		event.LocationId,
		event.LocationDetails,
		event.Price,
		event.CategoryId,
		event.AdditionalInfo,
		event.AdditionalUrl,
		event.ContactInfo,
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
func UpdateEvent(ctx context.Context, db *sql.DB, event *Event) error {
	if _, err := FindEventById(ctx, db, event.Id); err != nil {
		return err
	}

	query := `
		UPDATE events SET
			title = ?,
			description = ?,
			start_date = ?,
			end_date = ?,
			organization_id = ?,
			image_id = ?,
			location_id = ?,
			location_details = ?,
			price = ?,
			category_id = ?,
			additional_info = ?,
			additional_url = ?,
			contact_info = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	_, err := db.ExecContext(ctx, query,
		event.Title,
		event.Description,
		event.StartDate,
		event.EndDate,
		event.OrganizationId,
		event.ImageId,
		event.LocationId,
		event.LocationDetails,
		event.Price,
		event.CategoryId,
		event.AdditionalInfo,
		event.AdditionalUrl,
		event.ContactInfo,
		event.Id,
	)

	if err != nil {
		log.Printf("failed to update event: %s\n", err)
		return err
	}

	return nil
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
