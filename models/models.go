package models

import (
	"fmt"

	"github.com/somos831/somos-backend/errors/httperror"
)

// validationErr wraps a httperror.StatusError with a status code of 400.
func validationErr(msg string) error {
	err := httperror.BadRequest(msg)
	return fmt.Errorf("validation error: %w", err)
}
