package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
)

type ShortlinkRepositoryMock struct {
	mock.Mock
}

func (m *ShortlinkRepositoryMock) FindContentByShortlink(shortlink string) (*entities.Content, error) {
	args := m.Called(shortlink)
	return args.Get(0).(*entities.Content), args.Error(1)
}

func (m *ShortlinkRepositoryMock) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	args := m.Called(shortlink, text, expiryByMinutes)
	return args.Error(0)
}

func (m *ShortlinkRepositoryMock) DeleteExpiredContent() error {
	args := m.Called()
	return args.Error(0)
}
