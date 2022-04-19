package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zaenulhilmi/pastebin/services"
)

func NewShortlinkHandler(service services.ShortlinkService) PasteHandler {
	return &pasteHandler{
		shortlinkService: service,
	}
}

type PasteHandler interface {
	GetContent(w http.ResponseWriter, r *http.Request)
	CreateContent(w http.ResponseWriter, r *http.Request)
}

type pasteHandler struct {
	shortlinkService services.ShortlinkService
}

func (h *pasteHandler) GetContent(w http.ResponseWriter, r *http.Request) {
	shortlink := r.URL.Query().Get("shortlink")
	content, err := h.shortlinkService.GetContent(shortlink)
	if err != nil {
		internalServerError(w)
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

type CreateRequest struct {
	Text            string `json:"text"`
	ExpiryInMinutes int    `json:"expiry_in_minutes"`
}

func (h *pasteHandler) CreateContent(w http.ResponseWriter, r *http.Request) {
	var request CreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		internalServerError(w)
		return
	}
	shortlink, err := h.shortlinkService.CreateContent(request.Text, request.ExpiryInMinutes)
	if err != nil {
		internalServerError(w)
		return
	}
	w.Write([]byte("{\"shortlink\": \"" + shortlink + "\"}"))
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"error\": \"Something wrong\"}"))
}
