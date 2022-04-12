package services_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/services"
	"testing"
	"time"
)

func TestGetContentNotFound(t *testing.T) {
	var emptyContent *entities.Content
	repository := new(mocks.ShortlinkRepositoryMock)
	repository.On("FindContentByShortlink", "abc").Return(emptyContent, nil)

	shortlinkService := services.NewShortlinkService(repository)
	content, _ := shortlinkService.GetContent("abc")
	assert.Equal(t, emptyContent, content)
}

func TestGetContentOk(t *testing.T) {
	expectedContent := &entities.Content{
		Text:      "test",
		CreatedAt: time.Now(),
	}

	repository := new(mocks.ShortlinkRepositoryMock)
	repository.On("FindContentByShortlink", "abc").Return(expectedContent, nil)

	shortlinkService := services.NewShortlinkService(repository)
	content, _ := shortlinkService.GetContent("abc")
	assert.Equal(t, expectedContent, content)
}

func TestGetContentError(t *testing.T) {
	var emptyContent *entities.Content
	repository := new(mocks.ShortlinkRepositoryMock)
	repository.On("FindContentByShortlink", "abc").Return(emptyContent, errors.New("error"))

	shortlinkService := services.NewShortlinkService(repository)
	content, err := shortlinkService.GetContent("abc")
	assert.Equal(t, emptyContent, content)
	assert.Equal(t, "error", err.Error())
}
