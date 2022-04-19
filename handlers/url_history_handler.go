package handlers

import (
	"net/http"
)

func LoggingMiddleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		handler.ServeHTTP(w, r)
	})
}
