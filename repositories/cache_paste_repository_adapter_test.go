package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/repositories"
)

func TestFindContentByShortlinkAdapterByCache(t *testing.T) {
	shortlinkRepository := new(mocks.PasteRepositoryMock)
	shortlinkRepository.On("FindContentByShortlink", "shortlink").Return(&entities.Paste{Text: "from repo"}, nil)

	cache := new(mocks.CacheMock)
	cache.On("Get", "shortlink").Return(&entities.Paste{Text: "from cache"}, true)

	adapter := repositories.NewCacheAdapter(shortlinkRepository, cache)
	content, err := adapter.FindContentByShortlink("shortlink")
	assert.Nil(t, err)
	assert.Equal(t, content.Text, "from cache")
}

func TestFindContentByShortlinkAdapterByRepositories(t *testing.T) {
	shortlinkRepository := new(mocks.PasteRepositoryMock)
	shortlinkRepository.On("FindContentByShortlink", "shortlink").Return(&entities.Paste{Text: "from repo"}, nil)

	cache := new(mocks.CacheMock)
	var emptyContent *entities.Paste
	cache.On("Get", "shortlink").Return(emptyContent, false)

	adapter := repositories.NewCacheAdapter(shortlinkRepository, cache)
	content, err := adapter.FindContentByShortlink("shortlink")
	assert.Nil(t, err)
	assert.Equal(t, content.Text, "from repo")
}

func TestCreateContentSaveToCache(t *testing.T) {
	clock := new(mocks.ClockMock)
	shortlinkRepository := new(mocks.PasteRepositoryMock)
	shortlinkRepository.On("CreateContent", "shortlink", "from repo", 10).Return(nil)

	createdContent := &entities.Paste{Text: "from repo", CreatedAt: clock.Now(), ExpiryInMinutes: 10}
	shortlinkRepository.On("FindContentByShortlink", "shortlink").Return(createdContent, nil)

	cache := new(mocks.CacheMock)
	cache.On("Set", "shortlink", createdContent)

	adapter := repositories.NewCacheAdapter(shortlinkRepository, cache)
	err := adapter.CreateContent("shortlink", "from repo", 10)
	assert.Nil(t, err)
	shortlinkRepository.AssertCalled(t, "CreateContent", "shortlink", "from repo", 10)

	contentTextMatcher := mock.MatchedBy(func(content *entities.Paste) bool {
		return content.Text == "from repo"
	})
	cache.AssertCalled(t, "Set", "shortlink", contentTextMatcher)
}
