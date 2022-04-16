package handlers

import (
	"github.com/zaenulhilmi/pastebin/services"
	"net/http"
    "fmt"
)

func NewShortlinkHandler(service services.ShortlinkService) ShortlinkHandler {
	return &shortlinkHandler{
		shortlinkService: service,
	}
}

type ShortlinkHandler interface {
	GetContent(w http.ResponseWriter, r *http.Request)
}

type shortlinkHandler struct {
	shortlinkService services.ShortlinkService
}

func (h *shortlinkHandler) GetContent(w http.ResponseWriter, r *http.Request) {
	shortlink := r.URL.Query().Get("shortlink")

    fmt.Println(shortlink)
	content, err := h.shortlinkService.GetContent(shortlink)

    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("{\"error\": \"Something wrong\"}"))
        fmt.Println(err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
	if content == nil {
		w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("{\"error\": \"Shortlink not found\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	b, _ := content.MarshalJSON()
	w.Write(b)

}
