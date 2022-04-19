package repositories

import (
	"database/sql"

	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

type PasteRepository interface {
	FindContentByShortlink(shortlink string) (*entities.Content, error)
	CreateContent(shortlink string, text string, expiryByMinutes int) error
	DeleteExpiredContent() error
}

func NewShortlinkRepository(db *sql.DB, clock helpers.Clock) PasteRepository {
	return &pasteRepository{
		db:    db,
		clock: clock,
	}
}
