package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/url"
)

var ErrEventCategoryNotFound = errors.New("event category not found")

type EventCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// FindEventCategoryById finds a category in db using categoryId.
func FindEventCategoryById(ctx context.Context, db *sql.DB, categoryId int) (*EventCategory, error) {
	query := `SELECT * FROM event_categories WHERE id = ?`
	row := db.QueryRowContext(ctx, query, categoryId)

	var category EventCategory
	err := row.Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEventCategoryNotFound
		}
		log.Printf("failed to find event category by id:%s\nid: %d\n", err, categoryId)

		return nil, err
	}

	return &category, nil
}

// FindEventCategoryByName finds a category in db using name.
func FindEventCategoryByName(ctx context.Context, db *sql.DB, name string) (*EventCategory, error) {
	query := `SELECT * FROM event_categories WHERE name = ?`
	row := db.QueryRowContext(ctx, query, name)

	var category EventCategory
	err := row.Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEventCategoryNotFound
		}
		log.Printf("failed to find event category by name:%s\nname: %s\n", err, name)

		return nil, err
	}

	return &category, nil
}

// InsertEventCategory inserts into db. The id of the category is returned.
func InsertEventCategory(ctx context.Context, db *sql.DB, category EventCategory) (int, error) {
	query := `INSERT INTO event_categories (name) VALUES (?)`
	result, err := db.ExecContext(ctx, query, category.Name)
	if err != nil {
		log.Printf("failed to insert event category: %s\ncategory: %#v\n", err, category)
		return 0, err
	}

	categoryId, err := result.LastInsertId()
	if err != nil {
		log.Printf("failed to retreive event category id: %s\n", err)
		return 0, err
	}

	return int(categoryId), nil
}

// UpdateEventCategory updates category in db.
func UpdateEventCategory(ctx context.Context, db *sql.DB, category EventCategory) error {
	if _, err := FindEventCategoryById(ctx, db, category.Id); err != nil {
		return err
	}

	query := `UPDATE event_categories SET name = ? WHERE id = ?`
	_, err := db.ExecContext(ctx, query, category.Name, category.Id)
	if err != nil {
		log.Printf("failed to update event category: %s\ncategory: %#v\n", err, category)
		return err
	}

	return nil
}

// DeleteEventCategory deletes a category in db using categoryId.
func DeleteEventCategory(ctx context.Context, db *sql.DB, categoryId int) error {
	if _, err := FindEventCategoryById(ctx, db, categoryId); err != nil {
		return err
	}

	query := `DELETE FROM event_categories WHERE id = ?`
	_, err := db.ExecContext(ctx, query, categoryId)
	if err != nil {
		log.Printf("failed to delete event category: %s\nid: %d\n", err, categoryId)
		return err
	}

	return nil
}

// NewEventCategoryFromFormValues creates a new EventCategory from values.
func NewEventCategoryFromFormValues(values url.Values) (*EventCategory, error) {
	newCategory := &EventCategory{
		Name: values.Get("name"),
	}

	return newCategory, nil
}
