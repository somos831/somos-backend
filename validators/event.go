package validators

import (
	"context"
	"fmt"

	"github.com/somos831/somos-backend/models"
)

func (v *Validator) ValidateNewEvent(ctx context.Context, ev models.Event) error {
	errs := ValidationError{}

	if ev.Title == "" {
		errs.Add("title", "title cannot be empty")
	} else if len(ev.Title) > 100 {
		errs.Add("title", "title cannot be longer than 50 characters")
	}

	if ev.Description != nil && len(*ev.Description) > 1500 {
		errs.Add("description", "description cannot be longer than 1500 characters")
	}

	if ev.OrganizationId != nil {
		organizationExists, err := v.valuesExist(ctx, "organizations", "id", ev.OrganizationId)
		if err != nil {
			return err
		}

		if !organizationExists {
			errs.Add("organization_id", fmt.Sprintf("organization_id %d does not exist",
				*ev.OrganizationId))
		}
	}

	if ev.ImageId != nil {
		imageExists, err := v.valuesExist(ctx, "images", "id", ev.ImageId)
		if err != nil {
			return err
		}

		if !imageExists {
			errs.Add("image_id", fmt.Sprintf("image_id %d does not exist", *ev.ImageId))
		}
	}

	if ev.LocationId != nil {
		locationExists, err := v.valuesExist(ctx, "locations", "id", ev.LocationId)
		if err != nil {
			return err
		}

		if !locationExists {
			errs.Add("location_id", fmt.Sprintf("location_id %d does not exist",
				*ev.LocationId))
		}
	}

	if ev.LocationDetails != nil && len(*ev.LocationDetails) > 255 {
		errs.Add("location_details",
			"location_details cannot be longer than 255 characters")
	}

	if ev.CategoryId == 0 {
		errs.Add("category_id", "category_id cannot be empty")
	} else {
		categoryExists, err := v.valuesExist(ctx, "event_categories", "id", ev.CategoryId)
		if err != nil {
			return err
		}

		if !categoryExists {
			errs.Add("category_id", fmt.Sprintf("category_id %d does not exist",
				ev.CategoryId))
		}
	}

	if ev.AdditionalInfo != nil && len(*ev.AdditionalInfo) > 1500 {
		errs.Add("additional_info",
			"additional_info cannot be longer than 1500 characters")
	}

	if ev.AdditionalUrl != nil && len(*ev.AdditionalUrl) > 255 {
		errs.Add("additional_url",
			"additional_url cannot be longer than 255 characters")
	}

	if errs.None() {
		return nil
	}

	return errs
}
