package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
)

type ReadPasteRepositoryMock struct {
	mock.Mock
}

func (m *ReadPasteRepositoryMock) FindContentByShortlink(shortlink string) (*entities.Paste, error) {
	args := m.Called(shortlink)
	return args.Get(0).(*entities.Paste), args.Error(1)
}
