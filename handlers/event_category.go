package handlers

import (
	"errors"
	"net/http"
)

var errNonNumericEventCategoryId = errors.New("event category id must be an integer")

// ListAllCategories lists all the categories in the database.
func (s *Server) ListAllCategories(w http.ResponseWriter, r *http.Request) {
	panic("todo")
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
