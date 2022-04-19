package repositories

import (
	"database/sql"
	"fmt"

	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
)

type pasteRepository struct {
	db    *sql.DB
	clock helpers.Clock
}

const TABLE_NAME = "pastes"

func (s *pasteRepository) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	createdAt := s.clock.Now()
	query := fmt.Sprintf("INSERT INTO %s (shortlink, text, created_at, expiry_in_minutes) VALUES (?, ?, ?, ?)", TABLE_NAME)
	_, err := s.db.Exec(query, shortlink, text, createdAt, expiryByMinutes)
	if err != nil {
		return err
	}
	return nil
}

func (s *pasteRepository) FindContentByShortlink(shortlink string) (*entities.Paste, error) {
	var content entities.Paste
	query := fmt.Sprintf("SELECT text, created_at, expiry_in_minutes FROM %s WHERE shortlink = ?", TABLE_NAME)
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

func (s *pasteRepository) DeleteExpiredContent() error {
	query := fmt.Sprintf("DELETE FROM %s WHERE NOW() > DATE_ADD(created_at, INTERVAL expiry_in_minutes MINUTE) AND expiry_in_minutes != 0", TABLE_NAME)
	_, err := s.db.Exec(query)
	return err
}
