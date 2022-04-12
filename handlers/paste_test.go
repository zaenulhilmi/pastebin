package handlers_test

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
    "github.com/zaenulhilmi/pastebin/entities"
    "github.com/zaenulhilmi/pastebin/mocks"
    "github.com/zaenulhilmi/pastebin/handlers"
    "time"
)



func TestReadShortlinkOk(t *testing.T) {
    request, _ := http.NewRequest("GET", "http://localhost:8080/paste?shortlink=abc", nil)

    recorder := httptest.NewRecorder()

    shortlinkService := new(mocks.ShortlinkServiceMock)
    createdAt := time.Now()
    shortlinkService.On("GetContent", "abc").Return(&entities.Content{
        Text: "test",
        CreatedAt: createdAt,
        ExpiryInMinutes: 10,
    }, nil)

    shortlinkHandler  := handlers.NewShortlinkHandler(shortlinkService)

    handler := http.HandlerFunc(shortlinkHandler.GetContent)

    handler.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusOK, recorder.Code)
    assert.JSONEq(t, "{\"text\":\"test\",\"created_at\":\"" + createdAt.Format(time.RFC3339) + "\",\"expiry_in_minutes\":10}", recorder.Body.String())
}


func TestReadShortlinkNotFound(t *testing.T) {
    request, _ := http.NewRequest("GET", "http://localhost:8080/paste?shortlink=abc", nil)

    recorder := httptest.NewRecorder()

    shortlinkService := new(mocks.ShortlinkServiceMock)
    var emptyContent *entities.Content
    shortlinkService.On("GetContent", "abc").Return(emptyContent, nil)

    shortlinkHandler  := handlers.NewShortlinkHandler(shortlinkService)

    handler := http.HandlerFunc(shortlinkHandler.GetContent)

    handler.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusNotFound, recorder.Code)
}
