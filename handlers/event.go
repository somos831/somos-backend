package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/somos831/somos-backend/models"
	"github.com/somos831/somos-backend/responses"
)

var errNonNumericEventId = errors.New("event id must be an integer")

func (s *Server) ListEvents(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("limit")
	limit := 15

	if nStr != "" {
		n, err := strconv.Atoi(nStr)
		if err != nil {
			err = errors.Join(errNonNumericEventId, err)
			responses.Error(w, http.StatusBadRequest, err)

			return
		}

		limit = n
	}

	events, err := models.FindNRecentEvents(r.Context(), s.db, limit)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, events)
}

// GetEvent returns a single event by its id.
func (s *Server) GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventIdStr := params["id"]

	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		err = errors.Join(errNonNumericEventId, err)
		responses.Error(w, http.StatusBadRequest, err)

		return
	}

	event, err := models.FindEventById(r.Context(), s.db, eventId)
	if err != nil {
		responses.Error(w, http.StatusNotFound, err)
		return
	}

	responses.Json(w, http.StatusOK, event)
}

// CreateEvent creates a new event using the form data.
func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent models.Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = s.Validator.ValidateNewEvent(r.Context(), newEvent)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	eventId, err := models.InsertEvent(r.Context(), s.db, newEvent)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	res := map[string]int{
		"event_id": eventId,
	}
	responses.Json(w, http.StatusCreated, res)
}

// UpdateEvent updates an event by its id.
func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventIdStr := params["id"]

	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		err = errors.Join(errNonNumericEventId, err)
		responses.Error(w, http.StatusBadRequest, err)

		return
	}

	var event models.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	event.Id = eventId

	err = s.Validator.ValidateNewEvent(r.Context(), event)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = models.UpdateEvent(r.Context(), s.db, &event)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, event)
}

// DeleteEvent deletes an event by its id.
func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventIdStr := params["id"]

	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		err = errors.Join(errNonNumericEventId, err)
		responses.Error(w, http.StatusBadRequest, err)

		return
	}

	err = models.DeleteEvent(r.Context(), s.db, eventId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusNoContent, nil)
}
