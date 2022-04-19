package mocks

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type HandlerMock struct {
	mock.Mock
}

func (h *HandlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Called(w, r)
}
