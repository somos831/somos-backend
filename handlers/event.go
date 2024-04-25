package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/somos831/somos-backend/errors/httperror"
	"github.com/somos831/somos-backend/models"
)

type eventQueryParams struct {
	name        string
	description string
	location    string
	category    string
}

// ListAllEvents lists all the events in the database.
func (h *handler) ListAllEvents(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query()
	params := eventQueryParams{
		name:        query.Get("name"),
		description: query.Get("description"),
		location:    query.Get("location"),
		category:    query.Get("category"),
	}
	whereClause, whereArgs := buildWhereClause(params)

	results, err := h.db.Query(fmt.Sprintf(`
		SELECT events.id, events.name, events.description, events.location,
			categories.id as category_id, categories.name as category_name
		FROM events INNER JOIN categories
		ON events.category_id = categories.id
		%s;`, whereClause), whereArgs...)
	if err != nil {
		return httperror.InternalServer(err)
	}

	events := []models.Event{}
	for results.Next() {
		var event models.Event
		err = results.Scan(&event.Id, &event.Name, &event.Description,
			&event.Location, &event.Category.Id, &event.Category.Name)
		if err != nil {
			return httperror.InternalServer(err)
		}
		events = append(events, event)
	}

	if err = json.NewEncoder(w).Encode(events); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// GetEvent returns a single event by its id.
func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	id := r.PathValue("id")
	event, err := h.eventById(id)
	if err != nil {
		return err
	}
	if err = json.NewEncoder(w).Encode(event); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// CreateEvent creates a new event using the form data.
func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		return httperror.InternalServer(err)
	}

	var description *string
	if r.Form.Has("description") {
		desc := r.Form.Get("description")
		description = &desc
	}
	var location *string
	if r.Form.Has("location") {
		loc := r.Form.Get("location")
		location = &loc
	}
	event := models.Event{
		Name:        r.Form.Get("name"),
		Description: description,
		Location:    location,
	}
	if err := event.Validate(); err != nil {
		return err
	}

	category, err := h.CategoryById(r.Form.Get("categoryId"))
	if err != nil {
		return err
	}
	event.Category = *category

	result, err := h.db.Exec(`
		INSERT INTO events (name, description, category_id, location)
		VALUES (?, ?, ?, ?);
	`, event.Name, event.Description, event.Category.Id, event.Location)
	if err != nil {
		return httperror.InternalServer(err)
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		return httperror.InternalServer(err)
	}
	event.Id = int(eventId)

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(event); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// UpdateEvent updates an event by its id.
func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	event, err := h.eventById(r.PathValue("id"))
	if err != nil {
		return err
	}

	if err = r.ParseForm(); err != nil {
		return httperror.InternalServer(err)
	}
	if r.Form.Has("name") {
		event.Name = r.Form.Get("name")
	}
	if r.Form.Has("description") {
		desc := r.Form.Get("description")
		if desc == "" {
			event.Description = nil
		} else {
			event.Description = &desc
		}
	}
	if r.Form.Has("location") {
		loc := r.Form.Get("location")
		if loc == "" {
			event.Location = nil
		} else {
			event.Location = &loc
		}
	}
	if err := event.Validate(); err != nil {
		return err
	}

	if r.Form.Has("categoryId") {
		category, err := h.CategoryById(r.Form.Get("categoryId"))
		if err != nil {
			return err
		}
		event.Category = *category
	}

	_, err = h.db.Exec(`
		UPDATE events SET name = ?, description = ?, category_id = ?, location = ?
		WHERE id = ?;
	`, event.Name, event.Description, event.Category.Id, event.Location, event.Id)
	if err != nil {
		return httperror.InternalServer(err)
	}

	if err = json.NewEncoder(w).Encode(event); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// DeleteEvent deletes an event by its id.
func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	event, err := h.eventById(r.PathValue("id"))
	if err != nil {
		return err
	}

	_, err = h.db.Exec(`DELETE from events WHERE id = ?;`, event.Id)
	if err != nil {
		return httperror.InternalServer(err)
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

func buildWhereClause(params eventQueryParams) (string, []interface{}) {
	var whereClauseBuilder strings.Builder
	var whereArgs []interface{}

	if params.name != "" {
		whereClauseBuilder.WriteString("events.name LIKE ? AND ")
		whereArgs = append(whereArgs, "%"+params.name+"%")
	}
	if params.description != "" {
		whereClauseBuilder.WriteString("events.description LIKE ? AND ")
		whereArgs = append(whereArgs, "%"+params.description+"%")
	}
	if params.location != "" {
		whereClauseBuilder.WriteString("events.location LIKE ? AND ")
		whereArgs = append(whereArgs, "%"+params.location+"%")
	}
	if params.category != "" {
		category := params.category
		if categoryId, err := strconv.Atoi(category); err != nil {
			whereClauseBuilder.WriteString("categories.name LIKE ? AND ")
			whereArgs = append(whereArgs, "%"+category+"%")
		} else {
			whereClauseBuilder.WriteString("categories.id = ? AND ")
			whereArgs = append(whereArgs, categoryId)
		}
	}
	whereClause := strings.TrimSuffix(whereClauseBuilder.String(), " AND ")
	if len(whereClause) != 0 {
		whereClause = "WHERE " + whereClause
	}

	return whereClause, whereArgs
}

func (h *handler) eventById(id string) (*models.Event, error) {
	eventId, err := strconv.Atoi(id)
	if err != nil {
		return nil, httperror.BadRequest("id should be an integer")
	}
	result := h.db.QueryRow(`
		SELECT events.id, events.name, events.description, events.location,
			categories.id as category_id, categories.name as category_name
		FROM events INNER JOIN categories
		ON events.category_id = categories.id
		WHERE events.id = ?;
		`, eventId)

	var event models.Event
	err = result.Scan(&event.Id, &event.Name, &event.Description, &event.Location,
		&event.Category.Id, &event.Category.Name)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, httperror.NotFound(err)
	case err != nil:
		return nil, httperror.InternalServer(err)
	}

	return &event, nil
}
