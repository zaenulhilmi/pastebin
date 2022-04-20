package handlers

import (
	"net/http"

	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/services"
)

func LoggingMiddleware(logService services.LogService, handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logService.SaveLog(entities.Log{Url: "abc"})
		handler.ServeHTTP(w, r)
	})
}
