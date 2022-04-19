package repositories

import (
	"database/sql"

	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

type ReadPasteRepository interface {
	FindContentByShortlink(shortlink string) (*entities.Paste, error)
}

type WritePasteRepository interface {
	CreateContent(shortlink string, text string, expiryByMinutes int) error
	DeleteExpiredContent() error
}

func NewReadPasteRepository(db *sql.DB, clock helpers.Clock) ReadPasteRepository {
	return &readPasteRepository{
		db:    db,
		clock: clock,
	}
}

func NewWritePasteRepository(db *sql.DB, clock helpers.Clock) WritePasteRepository {
	return &writePasteRepository{
		db:    db,
		clock: clock,
	}
}
