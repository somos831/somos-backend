package models

import (
	"time"
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
		return validationErr("title cannot be empty")
	}
	if len(e.Title) > 100 {
		return validationErr("title must be less than 100 characters long")
	}
	if e.Description != nil && len(*e.Description) > 1500 {
		return validationErr("description must be less than 1500 characters long")
	}
	if e.Price < 0 {
		return validationErr("price cannot be negative")
	}

	return nil
}
