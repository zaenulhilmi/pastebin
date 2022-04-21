package handlers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/handlers"
	"github.com/zaenulhilmi/pastebin/mocks"
)

func TestRequestCreateMiddleware(t *testing.T) {
	testHandler := new(mocks.HandlerMock)
	logService := new(mocks.LogServiceMock)

	// param anything
	logService.On("SaveLog", mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/posts", nil)
	// add params
	req.URL.RawQuery = "shortlink=test"
	rr := httptest.NewRecorder()

	testHandler.On("ServeHTTP", rr, req)

	handler := handlers.LoggingMiddleware(logService, testHandler)
	handler.ServeHTTP(rr, req)
	testHandler.AssertCalled(t, "ServeHTTP", rr, req)
	logService.AssertCalled(t, "SaveLog", mock.Anything)
}

func TestRequestCreateMiddlewareNoShortlink(t *testing.T) {
	testHandler := new(mocks.HandlerMock)
	logService := new(mocks.LogServiceMock)

	// param anything
	logService.On("SaveLog", mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/posts", nil)
	// add params
	req.URL.RawQuery = "x=test"
	rr := httptest.NewRecorder()

	testHandler.On("ServeHTTP", rr, req)

	handler := handlers.LoggingMiddleware(logService, testHandler)
	handler.ServeHTTP(rr, req)
	testHandler.AssertCalled(t, "ServeHTTP", rr, req)
	logService.AssertNotCalled(t, "SaveLog", mock.Anything)
}
