package models

import (
	"time"
)

type Image struct {
	Id         int       `json:"id"`
	Filename   *string   `json:"filename"`
	Url        *string   `json:"url"`
	Alt        *string   `json:"alt"`
	UploadedAt time.Time `json:"uploadedAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (i Image) Validate() error {
	if i.Filename != nil && len(*i.Filename) > 255 {
		return validationErr("filename cannot be longer than 255 characters")
	}
	if i.Url != nil && len(*i.Url) > 255 {
		return validationErr("url cannot be longer than 255 characters")
	}
	if i.Alt != nil && len(*i.Alt) > 150 {
		return validationErr("alt cannot be longer than 150 characters")
	}

	return nil
}
