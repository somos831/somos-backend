package validators

import (
	"context"
	"errors"

	"github.com/somos831/somos-backend/models"
)

func (v *Validator) ValidateNewEvent(ctx context.Context, event models.Event) error {
	if event.Title == "" {
		return errors.New("event title is a required field")
	}

	if event.CategoryId == 0 {
		return errors.New("event category_id is a required field")
	}

	return nil
}
