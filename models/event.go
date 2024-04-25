package models

import "github.com/somos831/somos-backend/errors/httperror"

type Event struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Location    *string  `json:"location"`
	Category    Category `json:"category"`
}

func (e Event) Validate() error {
	if len(e.Name) == 0 {
		return httperror.BadRequest("name cannot be empty")
	}
	if len(e.Name) > 50 {
		return httperror.BadRequest("name cannot be longer than 50 characters")
	}
	if e.Description != nil && len(*e.Description) > 1000 {
		return httperror.BadRequest("description cannot be longer than 1000 characters")
	}
	if e.Location != nil && len(*e.Location) > 200 {
		return httperror.BadRequest("location cannot be longer than 200 characters")
	}

	return nil
}
