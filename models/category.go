package models

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c Category) Validate() error {
	if len(c.Name) == 0 {
		return validationErr("name cannot be empty")
	}
	if len(c.Name) > 50 {
		return validationErr("name cannot be longer than 50 characters")
	}

	return nil
}
