package models

// intoPtr creates a new pointer from val.
func intoPtr[T interface{}](val T) *T {
	return &val
}
