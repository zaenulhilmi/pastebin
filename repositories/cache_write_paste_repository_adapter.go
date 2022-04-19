package repositories

import (
	"github.com/zaenulhilmi/pastebin/helpers"
)

func NewWriteCacheAdapter(readRepository ReadPasteRepository, writeRepository WritePasteRepository, cache helpers.Cache) WritePasteRepository {
	return &cachePasteRepositoryAdapter{
		writeRepository: writeRepository,
		readRepository:  readRepository,
		cache:           cache,
	}
}

type cachePasteRepositoryAdapter struct {
	writeRepository WritePasteRepository
	readRepository  ReadPasteRepository
	cache           helpers.Cache
}

func (c *cachePasteRepositoryAdapter) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	err := c.writeRepository.CreateContent(shortlink, text, expiryByMinutes)
	if err != nil {
		return err
	}
	content, err := c.readRepository.FindContentByShortlink(shortlink)
	if err != nil {
		return err
	}
	c.cache.Set(shortlink, content)
	return err
}

func (c *cachePasteRepositoryAdapter) DeleteExpiredContent() error {
	err := c.writeRepository.DeleteExpiredContent()
	if err != nil {
		return err
	}
	return err
}
