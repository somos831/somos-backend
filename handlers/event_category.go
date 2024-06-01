package handlers

import (
	"errors"
	"net/http"

	"github.com/somos831/somos-backend/models"
	"github.com/somos831/somos-backend/responses"
)

var errNonNumericEventCategoryId = errors.New("event category id must be an integer")

// ListAllCategories lists all the categories in the database.
func (s *Server) ListAllCategories(w http.ResponseWriter, r *http.Request) {

	eventCategories, err := models.GetAllCategories(r.Context(), s.db)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, errors.New("failed to get event categories"))
	}

	responses.Json(w, http.StatusOK, eventCategories)
}

// GetCategory returns a single category by its id.
func (s *Server) GetCategory(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

// CreateCategory creates a new category using the form data.
func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

// UpdateCategory updates a category by its id.
func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

// DeleteCategory deletes a category by its id.
func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}
