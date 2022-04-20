package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
)

type LogServiceMock struct {
	mock.Mock
}

func (m *LogServiceMock) SaveLog(log entities.Log) error {
	args := m.Called(log)
	return args.Error(0)
}
