package models

import "github.com/somos831/somos-backend/errors/httperror"

type Location struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Url  *string `json:"url"`
}

func (l Location) Validate() error {
	if len(l.Name) > 255 {
		return httperror.BadRequest("name cannot be longer than 255 characters")
	}
	if l.Url != nil && len(*l.Url) > 255 {
		return httperror.BadRequest("url cannot be longer than 255 characters")
	}

	return nil
}
