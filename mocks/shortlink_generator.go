package mocks

import "github.com/stretchr/testify/mock"

type ShortlinkGeneratorMock struct {
	mock.Mock
}

func (s *ShortlinkGeneratorMock) Generate() string {
	args := s.Called()
	return args.String(0)
}
