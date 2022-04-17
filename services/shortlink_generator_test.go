package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/services"
)

func TestMd5Generator_GenerateOk(t *testing.T) {
	repo := new(mocks.ShortlinkRepositoryMock)
	var emptyContent *entities.Content
	repo.On("FindContentByShortlink", "abc").Return(emptyContent, nil)
	generator := services.NewShortlinkGenerator(repo, &MockToken{})
	shortlink, err := generator.Generate()
	assert.Nil(t, err)
	assert.NotEmpty(t, shortlink)
}

func TestMd5Generator_GenerateFail(t *testing.T) {
	repo := new(mocks.ShortlinkRepositoryMock)
	repo.On("FindContentByShortlink", "abc").Return(&entities.Content{Text: "something"}, nil)
	generator := services.NewShortlinkGenerator(repo, &MockToken{})
	shortlink, err := generator.Generate()
	assert.NotNil(t, err)
	assert.Empty(t, shortlink)
}

type MockToken struct{}

func (m *MockToken) Random(x int) string {
	return "abc"
}
