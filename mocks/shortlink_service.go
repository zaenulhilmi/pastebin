package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
)

type ShortlinkServiceMock struct {
	mock.Mock
}

func (m *ShortlinkServiceMock) GetContent(shortlink string) (*entities.Content, error) {
	args := m.Called(shortlink)
	return args.Get(0).(*entities.Content), args.Error(1)
}

func (m *ShortlinkServiceMock) CreateContent(text string, expiryInMinutes int) (string, error) {
	args := m.Called(text, expiryInMinutes)
	return args.String(0), args.Error(1)
}
