package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/somos831/somos-backend/models"
	"github.com/somos831/somos-backend/responses"
)

func (s *Server) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location models.Location

	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = s.Validator.NewLocation(location)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	locationID, err := models.InsertLocation(r.Context(), s.db, location)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	res := map[string]int{
		"location_id": locationID,
	}

	responses.Json(w, http.StatusCreated, res)
}
