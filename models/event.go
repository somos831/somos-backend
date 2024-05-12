package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/url"
	"strconv"
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
func UpdateEvent(ctx context.Context, db *sql.DB, event Event) error {
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
		return err
	}

	return nil
}

// EventExistsById checks if an event is in db using eventId.
func EventExistsById(ctx context.Context, db *sql.DB, eventId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM events WHERE id = ?`
	err := db.QueryRowContext(ctx, query, eventId).Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		log.Printf("failed to locate event by id: %s\nid: %d\n", err, eventId)

		return false, err
	}

	return count > 0, nil
}

// DeleteEvent deletes an event using eventId.
func DeleteEvent(ctx context.Context, db *sql.DB, eventId int) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.ExecContext(ctx, query, eventId)
	if err != nil {
		log.Printf("failed to delete event: %s\nid: %d\n", err, eventId)
		return err
	}

	return nil
}

// NewEventFromFormValues creates a new Event from values.
func NewEventFromFormValues(values url.Values) (*Event, error) {
	var newEvent Event

	if values.Has("organization_id") {
		organizationId, err := strconv.Atoi(values.Get("organization_id"))
		if err != nil {
			return nil, errors.Join(errors.New("organization_id should be a number"), err)
		}
		newEvent.OrganizationId = &organizationId
	}

	if values.Has("image_id") {
		imageId, err := strconv.Atoi(values.Get("image_id"))
		if err != nil {
			return nil, errors.Join(errors.New("imgae_id should be a number"), err)
		}
		newEvent.ImageId = &imageId
	}

	if values.Has("location_id") {
		locationId, err := strconv.Atoi(values.Get("location_id"))
		if err != nil {
			return nil, errors.Join(errors.New("location_id should be a number"), err)
		}
		newEvent.LocationId = &locationId
	}

	if values.Has("price") {
		price, err := strconv.ParseFloat(values.Get("price"), 32)
		if err != nil {
			return nil, errors.Join(errors.New("price should be a number"), err)
		}
		newEvent.Price = float32(price)
	}

	if values.Has("category_id") {
		categoryId, err := strconv.Atoi(values.Get("category_id"))
		if err != nil {
			return nil, errors.Join(errors.New("category_id should be a number"), err)
		}
		newEvent.CategoryId = categoryId
	}

	newEvent.Title = values.Get("title")
	newEvent.Description = intoPtr(values.Get("description"))
	newEvent.LocationDetails = intoPtr(values.Get("location_details"))
	newEvent.AdditionalInfo = intoPtr(values.Get("additional_info"))
	newEvent.AdditionalUrl = intoPtr(values.Get("additional_url"))

	return &newEvent, nil
}
