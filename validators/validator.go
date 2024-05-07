package validators

import "database/sql"

type Validator struct {
	DB *sql.DB
}

func NewValidator(db *sql.DB) Validator {
	return Validator{DB: db}
}
