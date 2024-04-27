package models

import (
	"time"

	"github.com/somos831/somos-backend/errors/httperror"
)

type Event struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	Organization string    `json:"organization"`
	ImgId        *int      `json:"imgId"`
	LocationId   *int      `json:"locationId"`
	Price        float32   `json:"price"`
	CategoryId   int       `json:"categoryId"`
	MoreInfoId   *int      `json:"moreInfoId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (e Event) Validate() error {
	if len(e.Title) == 0 {
		return httperror.BadRequest("title cannot be empty")
	}
	if len(e.Title) > 100 {
		return httperror.BadRequest("title must be less than 100 characters long")
	}
	if e.Description != nil && len(*e.Description) > 1500 {
		return httperror.BadRequest("description must be less than 1500 characters long")
	}
	if e.Price < 0 {
		return httperror.BadRequest("price cannot be negative")
	}

	return nil
}
