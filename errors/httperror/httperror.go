package httperror

import (
	"errors"
	"net/http"
)

type StatusError struct {
	error
	status int
}

func (e StatusError) Unwrap() error {
	return e.error
}

func (e StatusError) Status() int {
	return e.status
}

func InternalServer(err error) error {
	return StatusError{err, http.StatusInternalServerError}
}

func BadRequest(msg string) error {
	return StatusError{errors.New(msg), http.StatusBadRequest}
}

func NotFound(err error) error {
	return StatusError{err, http.StatusNotFound}
}
