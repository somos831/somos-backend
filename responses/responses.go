package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Json is responsible for returning the http responses
func Json(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// Error handles the error http responses
func Error(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Json(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})

		return
	}

	Json(w, http.StatusBadRequest, nil)
}
