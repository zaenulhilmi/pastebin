package repositories

import (
	"database/sql"
	"fmt"

	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

type readPasteRepository struct {
	db    *sql.DB
	clock helpers.Clock
}

const READ_TABLE_NAME = "pastes"

func (s *readPasteRepository) FindContentByShortlink(shortlink string) (*entities.Paste, error) {
	var content entities.Paste
	query := fmt.Sprintf("SELECT text, created_at, expiry_in_minutes FROM %s WHERE shortlink = ?", READ_TABLE_NAME)
	err := s.db.QueryRow(query, shortlink).
		Scan(&content.Text, &content.CreatedAt, &content.ExpiryInMinutes)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &content, nil
}
