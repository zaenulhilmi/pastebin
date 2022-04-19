package repositories

import (
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

func NewCacheAdapter(shortlinkRepository PasteRepository, cache helpers.Cache) PasteRepository {
	return &cachePasteRepositoryAdapter{
		shortlinkRepository: shortlinkRepository,
		cache:               cache,
	}
}

type cachePasteRepositoryAdapter struct {
	shortlinkRepository PasteRepository
	cache               helpers.Cache
}

func (c *cachePasteRepositoryAdapter) FindContentByShortlink(shortlink string) (*entities.Paste, error) {
	content, found := c.cache.Get(shortlink)
	if found {
		return content.(*entities.Paste), nil
	}
	return c.shortlinkRepository.FindContentByShortlink(shortlink)
}

func (c *cachePasteRepositoryAdapter) CreateContent(shortlink string, text string, expiryByMinutes int) error {
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

func (c *cachePasteRepositoryAdapter) DeleteExpiredContent() error {
	err := c.shortlinkRepository.DeleteExpiredContent()
	if err != nil {
		return err
	}
	return err
}
