package models

import "github.com/villaleo/somos-events/errors/httperror"

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c Category) Validate() error {
	if len(c.Name) == 0 {
		return httperror.BadRequest("name cannot be empty")
	}
	if len(c.Name) > 50 {
		return httperror.BadRequest("name cannot be longer than 50 characters")
	}

	return nil
}
