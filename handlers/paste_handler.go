package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zaenulhilmi/pastebin/services"
)

func NewShortlinkHandler(service services.PasteService) PasteHandler {
	return &pasteHandler{
		pasteService: service,
	}
}

type PasteHandler interface {
	GetContent(w http.ResponseWriter, r *http.Request)
	CreateContent(w http.ResponseWriter, r *http.Request)
}

type pasteHandler struct {
	pasteService services.PasteService
}

func (h *pasteHandler) GetContent(w http.ResponseWriter, r *http.Request) {
	shortlink := r.URL.Query().Get("shortlink")
	content, err := h.pasteService.GetContent(shortlink)
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
	shortlink, err := h.pasteService.CreateContent(request.Text, request.ExpiryInMinutes)
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
