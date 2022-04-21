package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
)

type LogRepositoryMock struct {
	mock.Mock
}

func (m *LogRepositoryMock) Create(log entities.ShortlinkLog) error {
	args := m.Called(log)
	return args.Error(0)
}
