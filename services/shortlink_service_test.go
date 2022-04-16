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

	generator := new(mocks.ShortlinkGeneratorMock)

	shortlinkService := services.NewShortlinkService(repository, generator)
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

	generator := new(mocks.ShortlinkGeneratorMock)
	shortlinkService := services.NewShortlinkService(repository, generator)
	content, _ := shortlinkService.GetContent("abc")
	assert.Equal(t, expectedContent, content)
}

func TestGetContentError(t *testing.T) {
	var emptyContent *entities.Content
	repository := new(mocks.ShortlinkRepositoryMock)
	repository.On("FindContentByShortlink", "abc").Return(emptyContent, errors.New("error"))

	generator := new(mocks.ShortlinkGeneratorMock)
	shortlinkService := services.NewShortlinkService(repository, generator)
	content, err := shortlinkService.GetContent("abc")
	assert.Equal(t, emptyContent, content)
	assert.Equal(t, "error", err.Error())
}

func TestCreateContentOk1(t *testing.T) {
	repository := new(mocks.ShortlinkRepositoryMock)
	expectShortlink := "abc"
	repository.On("CreateContent", expectShortlink, "content", 10).Return(nil)

	generator := new(mocks.ShortlinkGeneratorMock)
	generator.On("Generate").Return(expectShortlink)
	shortlinkService := services.NewShortlinkService(repository, generator)
	shortlink, err := shortlinkService.CreateContent("content", 10)
	assert.Nil(t, err)
	assert.Equal(t, expectShortlink, shortlink)
}

func TestCreateContentOk2(t *testing.T) {
	repository := new(mocks.ShortlinkRepositoryMock)
	expectShortlink := "xyz"
	repository.On("CreateContent", expectShortlink, "content", 10).Return(nil)

	generator := new(mocks.ShortlinkGeneratorMock)
	generator.On("Generate").Return(expectShortlink)
	shortlinkService := services.NewShortlinkService(repository, generator)
	shortlink, err := shortlinkService.CreateContent("content", 10)
	assert.Nil(t, err)
	assert.Equal(t, expectShortlink, shortlink)
}

func TestCreateContentError(t *testing.T) {
	repository := new(mocks.ShortlinkRepositoryMock)
	expectShortlink := "xyz"
	repository.On("generateShortlink").Return(expectShortlink)
	repository.On("CreateContent", expectShortlink, "content", 10).Return(errors.New("error"))

	generator := new(mocks.ShortlinkGeneratorMock)
	generator.On("Generate").Return(expectShortlink)
	shortlinkService := services.NewShortlinkService(repository, generator)
	shortlink, err := shortlinkService.CreateContent("content", 10)
	assert.NotNil(t, err)
	assert.Equal(t, "", shortlink)
}
