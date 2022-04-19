package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/repositories"
)

func TestFindContentByShortlinkAdapterByCache(t *testing.T) {
	shortlinkRepository := new(mocks.ReadPasteRepositoryMock)
	shortlinkRepository.On("FindContentByShortlink", "shortlink").Return(&entities.Paste{Text: "from repo"}, nil)

	cache := new(mocks.CacheMock)
	cache.On("Get", "shortlink").Return(&entities.Paste{Text: "from cache"}, true)

	adapter := repositories.NewCacheAdapter(shortlinkRepository, cache)
	content, err := adapter.FindContentByShortlink("shortlink")
	assert.Nil(t, err)
	assert.Equal(t, content.Text, "from cache")
}

func TestFindContentByShortlinkAdapterByRepositories(t *testing.T) {
	shortlinkRepository := new(mocks.ReadPasteRepositoryMock)
	shortlinkRepository.On("FindContentByShortlink", "shortlink").Return(&entities.Paste{Text: "from repo"}, nil)

	cache := new(mocks.CacheMock)
	var emptyContent *entities.Paste
	cache.On("Get", "shortlink").Return(emptyContent, false)

	adapter := repositories.NewCacheAdapter(shortlinkRepository, cache)
	content, err := adapter.FindContentByShortlink("shortlink")
	assert.Nil(t, err)
	assert.Equal(t, content.Text, "from repo")
}
