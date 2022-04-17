package mocks

import "time"

type ClockMock struct{}

func (m *ClockMock) Now() time.Time {
	return time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
}
