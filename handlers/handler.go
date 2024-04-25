package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}

type httpErrorResponse struct {
	Error string `json:"error"`
}

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error

func HttpHandler(f ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			log.Printf("%s\n", err)

			var status int
			var statusError interface {
				error
				Status() int
			}
			if errors.As(err, &statusError) {
				status = statusError.Status()
			}

			var message string
			switch status {
			case http.StatusBadRequest:
				message = err.Error()
			case http.StatusNotFound:
				message = "no results found"
			default:
				message = "an internal server error occurred"
			}

			res, _ := json.Marshal(httpErrorResponse{message})
			w.WriteHeader(status)
			if _, err = w.Write(res); err != nil {
				log.Printf("failed to write to ResponseWriter: %v\n", err)
			}
		}
	}
}
