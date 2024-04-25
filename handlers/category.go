package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/villaleo/somos-events/errors/httperror"
	"github.com/villaleo/somos-events/models"
)

// ListAllCategories lists all the categories in the database.
func (h *handler) ListAllCategories(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	var whereClause string
	whereArgs := []any{}
	if name := r.URL.Query().Get("name"); name != "" {
		whereClause = `WHERE name LIKE ?`
		whereArgs = append(whereArgs, "%"+name+"%")
	}

	query := fmt.Sprintf("SELECT * FROM categories %s;", whereClause)
	results, err := h.db.Query(query, whereArgs...)
	if err != nil {
		return httperror.InternalServer(err)
	}

	categories := []models.Category{}
	for results.Next() {
		var category models.Category
		err = results.Scan(&category.Id, &category.Name)
		if err != nil {
			return httperror.InternalServer(err)
		}
		categories = append(categories, category)
	}

	if err = json.NewEncoder(w).Encode(categories); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// GetCategory returns a single category by its id.
func (h *handler) GetCategory(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	id := r.PathValue("id")
	category, err := h.CategoryById(id)
	if err != nil {
		return err
	}
	if err = json.NewEncoder(w).Encode(category); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// CreateCategory creates a new category using the form data.
func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		return httperror.InternalServer(err)
	}
	category := models.Category{Name: r.Form.Get("name")}
	if err := category.Validate(); err != nil {
		return err
	}

	result, err := h.db.Exec(`
		INSERT INTO categories (name) VALUES (?);
	`, category.Name)
	if err != nil {
		return httperror.InternalServer(err)
	}
	categoryId, err := result.LastInsertId()
	if err != nil {
		return httperror.InternalServer(err)
	}
	category.Id = int(categoryId)

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(category); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// UpdateCategory updates a category by its id.
func (h *handler) UpdateCategory(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	id := r.PathValue("id")
	category, err := h.CategoryById(id)
	if err != nil {
		return err
	}

	if err := r.ParseForm(); err != nil {
		return httperror.InternalServer(err)
	}
	updatedCategory := models.Category{Name: r.Form.Get("name")}
	if err := updatedCategory.Validate(); err != nil {
		return err
	}

	query := `UPDATE categories SET name = ? WHERE id = ?;`
	_, err = h.db.Exec(query, updatedCategory.Name, category.Id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return httperror.NotFound(err)
	case err != nil:
		return httperror.InternalServer(err)
	}

	updatedCategory.Id = category.Id
	if err = json.NewEncoder(w).Encode(updatedCategory); err != nil {
		return httperror.InternalServer(err)
	}

	return nil
}

// DeleteCategory deletes a category by its id.
func (h *handler) DeleteCategory(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	id := r.PathValue("id")
	category, err := h.CategoryById(id)
	if err != nil {
		return err
	}

	query := `DELETE from categories WHERE id = ?;`
	_, err = h.db.Exec(query, category.Id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return httperror.NotFound(err)
	case err != nil:
		return httperror.InternalServer(err)
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

// CategoryById returns a Category corresponding to the id, if it exists.
func (h *handler) CategoryById(id string) (*models.Category, error) {
	if len(id) == 0 {
		return nil, httperror.BadRequest("name cannot be empty")
	}
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return nil, httperror.BadRequest("category id should be an integer")
	}
	result := h.db.QueryRow(`SELECT * FROM categories WHERE id = ?`, categoryId)

	var category models.Category
	err = result.Scan(&category.Id, &category.Name)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, httperror.NotFound(err)
	case err != nil:
		return nil, httperror.InternalServer(err)
	}

	return &category, nil
}
