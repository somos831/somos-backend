package models

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password,omitempty"` // TODO: change to hashed when auth is implemented
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	StatusID       int    `json:"status_id"`
	RoleID         int    `json:"role_id"`
}

func FindUserByID(db *sql.DB, userID int) (*User, error) {

	row := db.QueryRow(`SELECT id, username, email, first_name, last_name, profile_picture, status_id, role_id FROM users WHERE id = ?`, userID)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.ProfilePicture,
		&user.StatusID,
		&user.RoleID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

// TODO: will need to modify once we add authorization
func InsertUser(db *sql.DB, user *User) (int, error) {

	query := "INSERT INTO users (username, email, password, first_name, last_name, profile_picture, status_id, role_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := db.Exec(query, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.ProfilePicture, user.StatusID, user.RoleID)
	if err != nil {
		log.Printf("Failed to create user due to: %s\n", err.Error())
		return 0, err
	}

	// Get the ID of the newly inserted user
	userID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to fetch created user ID: %s\n", err.Error())
		return 0, err
	}

	return int(userID), nil
}

func UpdateUser(db *sql.DB, user *User) error {

	query := "UPDATE users SET username=?, email=?, first_name=?, last_name=?, profile_picture=?, status_id=?, role_id=?, updated_at = CURRENT_TIMESTAMP WHERE id=?"

	_, err := db.Exec(query, user.Username, user.Email, user.FirstName, user.LastName, user.ProfilePicture, user.StatusID, user.RoleID, user.ID)
	if err != nil {
		log.Printf("Failed to update user due to: %s", err.Error())
		return err
	}
	return nil
}

func UserExistsByEmail(db *sql.DB, email string) (bool, error) {

	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func UserExistsByUsername(db *sql.DB, username string) (bool, error) {

	var count int
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func DeleteUser(db *sql.DB, userID int) error {

	_, err := db.Exec("DELETE FROM users WHERE ID = ?", userID)
	if err != nil {
		log.Printf("Failed to delete user due to: %s\n", err.Error())
		return err
	}

	return nil
}
