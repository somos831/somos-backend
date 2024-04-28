package models

type Organization struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (p Organization) Validate() error {
	if len(p.Name) > 100 {
		return validationErr("name must be less than 100 characters long")
	}

	return nil
}
