package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/repositories"
)

func TestCreateContentSaveToCache(t *testing.T) {
	clock := new(mocks.ClockMock)
	writePasteRepository := new(mocks.WritePasteRepositoryMock)
	writePasteRepository.On("CreateContent", "shortlink", "from repo", 10).Return(nil)

	readPasteRepository := new(mocks.ReadPasteRepositoryMock)
	createdContent := &entities.Paste{Text: "from repo", CreatedAt: clock.Now(), ExpiryInMinutes: 10}
	readPasteRepository.On("FindContentByShortlink", "shortlink").Return(createdContent, nil)

	cache := new(mocks.CacheMock)
	cache.On("Set", "shortlink", createdContent)

	adapter := repositories.NewWriteCacheAdapter(readPasteRepository, writePasteRepository, cache)
	err := adapter.CreateContent("shortlink", "from repo", 10)
	assert.Nil(t, err)
	writePasteRepository.AssertCalled(t, "CreateContent", "shortlink", "from repo", 10)

	contentTextMatcher := mock.MatchedBy(func(content *entities.Paste) bool {
		return content.Text == "from repo"
	})
	cache.AssertCalled(t, "Set", "shortlink", contentTextMatcher)
}
