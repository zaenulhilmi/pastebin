package handlers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/zaenulhilmi/pastebin/handlers"
	"github.com/zaenulhilmi/pastebin/mocks"
)

func TestRequestLogger(t *testing.T) {
	testHandler := new(mocks.HandlerMock)
	req := httptest.NewRequest("GET", "/posts", nil)
	rr := httptest.NewRecorder()

	testHandler.On("ServeHTTP", rr, req)

	handler := handlers.LoggingMiddleware(testHandler)
	handler.ServeHTTP(rr, req)
	testHandler.AssertCalled(t, "ServeHTTP", rr, req)
}
