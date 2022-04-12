package handlers_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.PingHandler)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
