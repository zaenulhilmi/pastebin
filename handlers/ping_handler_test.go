package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/handlers"
)

func TestPingHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.PingHandler)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
