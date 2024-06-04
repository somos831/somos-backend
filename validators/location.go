package validators

import "github.com/somos831/somos-backend/models"

func (v *Validator) NewLocation(location models.Location) error {
	errs := ValidationError{}

	if location.Name == "" {
		errs.Add("name", "location name cannot be empty")
	}

	if location.Address == "" {
		errs.Add("address", "location address cannot be empty")
	}

	if errs.None() {
		return nil
	}

	return errs
}
