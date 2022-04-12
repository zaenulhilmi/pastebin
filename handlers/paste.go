package handlers

import (
    "net/http"
    "github.com/zaenulhilmi/pastebin/services"
    "encoding/json"
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

    content, _ := h.shortlinkService.GetContent(shortlink)

    if content == nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    // marshal content to json 
    b, _ := json.Marshal(content)
    w.Write(b)

}

