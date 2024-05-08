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

var (
	UserNotFoundErr = errors.New("user not found")
)

// Create User: Endpoint for creating a new user account.
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	// Parse JSON request body into a User struct
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("Failed to decode request body"))
		return
	}

	err = s.Validator.ValidateNewUser(r.Context(), newUser)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Insert the new user into the database
	userID, err := models.InsertUser(r.Context(), s.db, &newUser)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, errors.New("Failed to create user"))
		return
	}

	// Return the ID of the newly created user in the response
	jsonResponse := map[string]int{"userID": userID}
	responses.Json(w, http.StatusCreated, jsonResponse)
}

// Retrieve User: Endpoint for retrieving user information.
func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	userID, err := strconv.Atoi(id)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("User ID must be an integer value"))
		return
	}

	foundUser, err := models.FindUserByID(r.Context(), s.db, userID)

	if errors.Is(err, UserNotFoundErr) {
		responses.Error(w, http.StatusNotFound, errors.New("User with given ID does not exist"))
		return
	}
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, errors.New("Failed to get user"))
		return
	}

	responses.Json(w, http.StatusFound, foundUser)
}

// UpdateUser: Endpoint for updating user information.
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	userID, err := strconv.Atoi(id)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("User ID must be an integer value"))
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	user.ID = userID

	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("Failed to decode request body"))
		return
	}

	err = s.Validator.ValidateUpdatedFields(r.Context(), user)
	if errors.Is(err, UserNotFoundErr) {
		responses.Error(w, http.StatusNotFound, errors.New("User with given ID does not exist"))
		return
	}
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = models.UpdateUser(r.Context(), s.db, &user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, errors.New("Failed to update user"))
		return
	}

	responses.Json(w, http.StatusOK, user)
}

// Delete User: Endpoint for deleting a user account.
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	userID, err := strconv.Atoi(id)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("User ID must be an integer value"))
		return
	}

	user, err := models.FindUserByID(r.Context(), s.db, userID)

	if errors.Is(err, UserNotFoundErr) {
		responses.Error(w, http.StatusNotFound, errors.New("User with given ID was not found"))
		return
	}
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, errors.New("Failed to delete user"))
		return
	}

	err = models.DeleteUser(r.Context(), s.db, user.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, errors.New("Failed to delete user"))
		return
	}

	responses.Json(w, http.StatusNoContent, nil)
}
