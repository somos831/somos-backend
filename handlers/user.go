package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/somos831/somos-backend/models"
	"github.com/somos831/somos-backend/responses"
)

// Create User: Endpoint for creating a new user account.
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	// Parse JSON request body into a User struct
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Failed to decode request body"))
		return
	}

	err = s.Validator.ValidateNewUser(newUser)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Insert the new user into the database
	userID, err := models.InsertUser(s.db, &newUser)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Failed to create user"))
		return
	}

	// Return the ID of the newly created user in the response
	jsonResponse := map[string]int{"userID": userID}
	responses.JSON(w, http.StatusCreated, jsonResponse)
}

// Retrieve User: Endpoint for retrieving user information.
func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("User ID must be an integer value"))
		return
	}

	foundUser, err := models.FindUserByID(s.db, userID)

	if err != nil {
		if err.Error() == "user not found" {
			responses.ERROR(w, http.StatusNotFound, errors.New("User with given ID does not exist"))
		} else {
			responses.ERROR(w, http.StatusInternalServerError, errors.New("Failed to get user"))
		}
		return
	}

	responses.JSON(w, http.StatusFound, foundUser)
}

// UpdateUser: Endpoint for updating user information.
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("User ID must be an integer value"))
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	user.ID = userID

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Failed to decode request body"))
		return
	}

	err = s.Validator.ValidateUpdatedFields(user)
	if err != nil {

		if err.Error() == "user not found" {
			responses.ERROR(w, http.StatusNotFound, errors.New("User with given ID does not exist"))
		} else {
			responses.ERROR(w, http.StatusBadRequest, err)
		}
		return
	}

	err = models.UpdateUser(s.db, &user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Failed to update user"))
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// Delete User: Endpoint for deleting a user account.
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("User ID must be an integer value"))
		return
	}

	user, err := models.FindUserByID(s.db, userID)

	if err != nil {
		if err.Error() == "user not found" {
			responses.ERROR(w, http.StatusNotFound, errors.New("User with given ID was not found"))
			return
		} else {
			responses.ERROR(w, http.StatusInternalServerError, errors.New("Failed to delete user"))
			return
		}
	}

	err = models.DeleteUser(s.db, user.ID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Failed to delete user"))
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
