package handlers

import (
	"context"
	"net/http"
	"time"
)

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context for the request
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		// Pass the context to the next handler in the chain
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
