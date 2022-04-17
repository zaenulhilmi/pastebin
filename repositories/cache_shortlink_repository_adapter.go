package repositories

import (
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

func NewCacheAdapter(shortlinkRepository ShortlinkRepository, cache helpers.Cache) ShortlinkRepository {
	return &cacheShortlinkRepositoryAdapter{
		shortlinkRepository: shortlinkRepository,
		cache:               cache,
	}
}

type cacheShortlinkRepositoryAdapter struct {
	shortlinkRepository ShortlinkRepository
	cache               helpers.Cache
}

func (c *cacheShortlinkRepositoryAdapter) FindContentByShortlink(shortlink string) (*entities.Content, error) {
	content, found := c.cache.Get(shortlink)
	if found {
		return content.(*entities.Content), nil
	}
	return c.shortlinkRepository.FindContentByShortlink(shortlink)
}

func (c *cacheShortlinkRepositoryAdapter) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	err := c.shortlinkRepository.CreateContent(shortlink, text, expiryByMinutes)
	if err != nil {
		return err
	}
	content, err := c.shortlinkRepository.FindContentByShortlink(shortlink)
	if err != nil {
		return err
	}
	c.cache.Set(shortlink, content)
	return err
}
