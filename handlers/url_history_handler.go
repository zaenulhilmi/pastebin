package handlers

import (
	"net/http"
	"time"

	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/services"
)

func LoggingMiddleware(logService services.LogService, handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shortlink := r.URL.Query().Get("shortlink")

		if shortlink != "" {
			log := entities.ShortlinkLog{
				Method:    r.Method,
				Shortlink: r.URL.RawQuery,
				Address:   r.RemoteAddr,
				CreatedAt: time.Now(),
			}

			logService.SaveLog(log)
		}
		handler.ServeHTTP(w, r)
	})
}
