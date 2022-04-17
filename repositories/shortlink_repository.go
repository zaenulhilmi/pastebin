package repositories

import (
	"database/sql"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

type ShortlinkRepository interface {
	FindContentByShortlink(shortlink string) (*entities.Content, error)
	CreateContent(shortlink string, text string, expiryByMinutes int) error
}

func NewShortlinkRepository(db *sql.DB, clock helpers.Clock) ShortlinkRepository {
	return &shortlinkRepository{
		db:    db,
		clock: clock,
	}
}


