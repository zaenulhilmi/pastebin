package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
)

type PasteServiceMock struct {
	mock.Mock
}

func (m *PasteServiceMock) GetContent(shortlink string) (*entities.Paste, error) {
	args := m.Called(shortlink)
	return args.Get(0).(*entities.Paste), args.Error(1)
}

func (m *PasteServiceMock) CreateContent(text string, expiryInMinutes int) (string, error) {
	args := m.Called(text, expiryInMinutes)
	return args.String(0), args.Error(1)
}

func (m *PasteServiceMock) DeleteExpiredContent() error {
	args := m.Called()
	return args.Error(0)
}
