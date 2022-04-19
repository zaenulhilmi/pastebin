package mocks

import (
	"github.com/stretchr/testify/mock"
)

type WritePasteRepositoryMock struct {
	mock.Mock
}

func (m *WritePasteRepositoryMock) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	args := m.Called(shortlink, text, expiryByMinutes)
	return args.Error(0)
}

func (m *WritePasteRepositoryMock) DeleteExpiredContent() error {
	args := m.Called()
	return args.Error(0)
}
