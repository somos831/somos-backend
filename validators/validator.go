package validators

import (
	"context"
	"database/sql"
	"fmt"
)

type Validator struct {
	DB *sql.DB
}

func NewValidator(db *sql.DB) Validator {
	return Validator{DB: db}
}

// ValidationError maps fields to reason(s) why the validation failed for that
// field. A ValidationError may be empty, but not nil.
//
//	errs := ValidationError{}
//
//	if errs.None() {
//	   return nil
//	}
//
// Use None() to check if there are any validation errors.
type ValidationError map[string][]string

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: %#v", v)
}

// Add adds a new validation error to the collection.
func (v *ValidationError) Add(key string, val string) {
	(*v)[key] = append((*v)[key], val)
}

// None returns true if there are no validation errors.
func (v ValidationError) None() bool {
	return len(v) == 0
}

// valuesExist queries the database to check if any values exist
// in a table using column col and condition value val.
//
// SELECT COUNT(*) FROM table WHERE col = val
func (m *Validator) valuesExist(ctx context.Context, table, col string, val interface{}) (bool, error) {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE %s = ?`, table, col)
	row := m.DB.QueryRowContext(ctx, query, val)

	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}
