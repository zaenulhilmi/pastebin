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
