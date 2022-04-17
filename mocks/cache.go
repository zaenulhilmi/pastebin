package mocks

import "github.com/stretchr/testify/mock"

type CacheMock struct {
	mock.Mock
}

func (c *CacheMock) Get(key string) (interface{}, bool) {
	args := c.Called(key)
	return args.Get(0), args.Bool(1)
}

func (c *CacheMock) Set(key string, value interface{}) {
	c.Called(key, value)
}

func (c *CacheMock) Delete(key string) {
	c.Called(key)
}
