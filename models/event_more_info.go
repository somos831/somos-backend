package models

type EventMoreInfo struct {
	Id   int     `json:"id"`
	Info *string `json:"info"`
	Url  *string `json:"url"`
}

func (e EventMoreInfo) Validate() error {
	if e.Info != nil && len(*e.Info) > 1500 {
		return validationErr("info must be less than 1500 characters long")
	}
	if e.Url != nil && len(*e.Url) > 255 {
		return validationErr("url must be less than 255 characters long")
	}

	return nil
}
