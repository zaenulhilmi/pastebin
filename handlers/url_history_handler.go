package handlers

import (
	"net/http"

	"github.com/zaenulhilmi/pastebin/entities"
)

type LogService interface {
	SaveLog(entities.Log) error
}

func LoggingMiddleware(logService LogService, handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here

		logService.SaveLog(entities.Log{Url: "abc"})
		handler.ServeHTTP(w, r)
	})
}
