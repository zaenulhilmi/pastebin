package repositories

import (
	"database/sql"

	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

type pasteRepository struct {
	db    *sql.DB
	clock helpers.Clock
}

func (s *pasteRepository) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	createdAt := s.clock.Now()
	_, err := s.db.Exec("INSERT INTO contents (shortlink, text, created_at, expiry_in_minutes) VALUES (?, ?, ?, ?)", shortlink, text, createdAt, expiryByMinutes)
	if err != nil {
		return err
	}
	return nil
}

func (s *pasteRepository) FindContentByShortlink(shortlink string) (*entities.Paste, error) {
	var content entities.Paste
	err := s.db.QueryRow("SELECT text, created_at, expiry_in_minutes FROM contents WHERE shortlink = ?", shortlink).
		Scan(&content.Text, &content.CreatedAt, &content.ExpiryInMinutes)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (s *pasteRepository) DeleteExpiredContent() error {
	_, err := s.db.Exec("DELETE FROM contents WHERE NOW() > DATE_ADD(created_at, INTERVAL expiry_in_minutes MINUTE) AND expiry_in_minutes != 0")
	return err
}
