package handlers_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/handlers"
	"github.com/zaenulhilmi/pastebin/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestReadShortlinkOk(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/paste?shortlink=abc", nil)

	recorder := httptest.NewRecorder()

	shortlinkService := new(mocks.ShortlinkServiceMock)
	createdAt := time.Now()
	shortlinkService.On("GetContent", "abc").Return(&entities.Content{
		Text:            "test",
		CreatedAt:       createdAt,
		ExpiryInMinutes: 10,
	}, nil)

	shortlinkHandler := handlers.NewShortlinkHandler(shortlinkService)

	handler := http.HandlerFunc(shortlinkHandler.GetContent)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, "{\"text\":\"test\",\"created_at\":\""+createdAt.Format(time.RFC3339)+"\",\"expiry_in_minutes\":10}", recorder.Body.String())
}

func TestReadShortlinkNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/paste?shortlink=abc", nil)

	recorder := httptest.NewRecorder()

	shortlinkService := new(mocks.ShortlinkServiceMock)
	var emptyContent *entities.Content
	shortlinkService.On("GetContent", "abc").Return(emptyContent, nil)

	shortlinkHandler := handlers.NewShortlinkHandler(shortlinkService)

	handler := http.HandlerFunc(shortlinkHandler.GetContent)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.JSONEq(t, "{\"error\":\"Shortlink not found\"}", recorder.Body.String())
}

func TestReadShortlinkGeneralError(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/paste?shortlink=abc", nil)

	recorder := httptest.NewRecorder()

	shortlinkService := new(mocks.ShortlinkServiceMock)
	var emptyContent *entities.Content
	shortlinkService.On("GetContent", "abc").Return(emptyContent, errors.New("error"))

	shortlinkHandler := handlers.NewShortlinkHandler(shortlinkService)

	handler := http.HandlerFunc(shortlinkHandler.GetContent)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.JSONEq(t, "{\"error\":\"Something wrong\"}", recorder.Body.String())
}

func TestCreateShortlinkContent(t *testing.T) {
	param := "test"
	request, _ := http.NewRequest("POST", "http://localhost:8080/paste",
		strings.NewReader("{\"text\":\""+param+"\",\"expiry_in_minutes\":10}"))

	recorder := httptest.NewRecorder()

	shortlinkService := new(mocks.ShortlinkServiceMock)
	shortlinkService.On("CreateContent", param, 10).Return("abc", nil)

	shortlinkHandler := handlers.NewShortlinkHandler(shortlinkService)

	handler := http.HandlerFunc(shortlinkHandler.CreateContent)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, "{\"shortlink\":\"abc\"}", recorder.Body.String())
}

func TestCreateShortlinkContent2(t *testing.T) {

	param := "test1"
	request, _ := http.NewRequest("POST", "http://localhost:8080/paste",
		strings.NewReader("{\"text\":\""+param+"\",\"expiry_in_minutes\":10}"))

	recorder := httptest.NewRecorder()

	shortlinkService := new(mocks.ShortlinkServiceMock)
	shortlinkService.On("CreateContent", param, 10).Return("xyz", nil)

	shortlinkHandler := handlers.NewShortlinkHandler(shortlinkService)

	handler := http.HandlerFunc(shortlinkHandler.CreateContent)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, "{\"shortlink\":\"xyz\"}", recorder.Body.String())
}

func TestCreateShortlinkUnknownError(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://localhost:8080/paste",
		strings.NewReader("{\"text\":\"test\",\"expiry_in_minutes\":10}"))
	recorder := httptest.NewRecorder()

	shortlinkService := new(mocks.ShortlinkServiceMock)
	shortlinkService.On("CreateContent", "test", 10).Return("", errors.New("error"))

	shortlinkHandler := handlers.NewShortlinkHandler(shortlinkService)

	handler := http.HandlerFunc(shortlinkHandler.CreateContent)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.JSONEq(t, "{\"error\":\"Something wrong\"}", recorder.Body.String())
}
