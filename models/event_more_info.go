package models

import "github.com/somos831/somos-backend/errors/httperror"

type EventMoreInfo struct {
	Id   int     `json:"id"`
	Info *string `json:"info"`
	Url  *string `json:"url"`
}

func (e EventMoreInfo) Validate() error {
	if e.Info != nil && len(*e.Info) > 1500 {
		return httperror.BadRequest("info cannot be longer than 1500 characters")
	}
	if e.Url != nil && len(*e.Url) > 255 {
		return httperror.BadRequest("url cannot be longer than 255 characters")
	}

	return nil
}
