package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/services"
)

func TestGetContentNotFound(t *testing.T) {
	var emptyContent *entities.Paste
	readRepository := new(mocks.ReadPasteRepositoryMock)
	readRepository.On("FindContentByShortlink", "abc").Return(emptyContent, nil)
	writeRepository := new(mocks.WritePasteRepositoryMock)

	generator := new(mocks.ShortlinkGeneratorMock)

	shortlinkService := services.NewShortlinkService(readRepository, writeRepository, generator)
	content, _ := shortlinkService.GetContent("abc")
	assert.Equal(t, emptyContent, content)
}

func TestGetContentOk(t *testing.T) {
	expectedContent := &entities.Paste{
		Text:      "test",
		CreatedAt: time.Now(),
	}

	readRepository := new(mocks.ReadPasteRepositoryMock)
	readRepository.On("FindContentByShortlink", "abc").Return(expectedContent, nil)
	writeRepository := new(mocks.WritePasteRepositoryMock)

	generator := new(mocks.ShortlinkGeneratorMock)
	shortlinkService := services.NewShortlinkService(readRepository, writeRepository, generator)
	content, _ := shortlinkService.GetContent("abc")
	assert.Equal(t, expectedContent, content)

}

func TestGetContentError(t *testing.T) {
	var emptyContent *entities.Paste
	readRepository := new(mocks.ReadPasteRepositoryMock)
	readRepository.On("FindContentByShortlink", "abc").Return(emptyContent, errors.New("error"))
	writeRepository := new(mocks.WritePasteRepositoryMock)

	generator := new(mocks.ShortlinkGeneratorMock)
	shortlinkService := services.NewShortlinkService(readRepository, writeRepository, generator)
	content, err := shortlinkService.GetContent("abc")
	assert.Equal(t, emptyContent, content)
	assert.Equal(t, "error", err.Error())
}

func TestCreateContentOk(t *testing.T) {
	tests := []struct {
		name            string
		expectShortlink string
	}{
		{
			name:            "abc",
			expectShortlink: "abc",
		},
		{
			name:            "xyz",
			expectShortlink: "xyz",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writeRepository := new(mocks.WritePasteRepositoryMock)
			writeRepository.On("CreateContent", test.expectShortlink, "content", 10).Return(nil)
			readRepository := new(mocks.ReadPasteRepositoryMock)

			generator := new(mocks.ShortlinkGeneratorMock)
			generator.On("Generate").Return(test.expectShortlink, nil)

			shortlinkService := services.NewShortlinkService(readRepository, writeRepository, generator)
			shortlink, err := shortlinkService.CreateContent("content", 10)

			assert.Nil(t, err)
			assert.Equal(t, test.expectShortlink, shortlink)

		})
	}
}

func TestCreateContentError(t *testing.T) {
	writeRepository := new(mocks.WritePasteRepositoryMock)
	expectShortlink := "xyz"
	writeRepository.On("generateShortlink").Return(expectShortlink)
	writeRepository.On("CreateContent", expectShortlink, "content", 10).Return(errors.New("error"))
	readRepository := new(mocks.ReadPasteRepositoryMock)

	generator := new(mocks.ShortlinkGeneratorMock)
	generator.On("Generate").Return(expectShortlink, nil)
	shortlinkService := services.NewShortlinkService(readRepository, writeRepository, generator)
	shortlink, err := shortlinkService.CreateContent("content", 10)
	assert.NotNil(t, err)
	assert.Equal(t, "", shortlink)
}

func TestDeleteExpiredContent(t *testing.T) {
	writeRepository := new(mocks.WritePasteRepositoryMock)
	writeRepository.On("DeleteExpiredContent").Return(nil)
	readRepository := new(mocks.ReadPasteRepositoryMock)

	shortlinkService := services.NewShortlinkService(readRepository, writeRepository, nil)
	err := shortlinkService.DeleteExpiredContent()
	assert.Nil(t, err)
	writeRepository.AssertCalled(t, "DeleteExpiredContent")
}

func TestDeleteExpiredContentError(t *testing.T) {
	writeRepository := new(mocks.WritePasteRepositoryMock)
	writeRepository.On("DeleteExpiredContent").Return(errors.New("error"))
	readRepository := new(mocks.ReadPasteRepositoryMock)

	shortlinkService := services.NewShortlinkService(readRepository, writeRepository, nil)
	err := shortlinkService.DeleteExpiredContent()
	assert.NotNil(t, err)
	writeRepository.AssertCalled(t, "DeleteExpiredContent")
}
