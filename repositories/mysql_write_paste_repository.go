package repositories

import (
	"database/sql"
	"fmt"

	"github.com/zaenulhilmi/pastebin/helpers"
)

type writePasteRepository struct {
	db    *sql.DB
	clock helpers.Clock
}

const WRITE_TABLE_NAME = "pastes"

func (s *writePasteRepository) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	createdAt := s.clock.Now()
	query := fmt.Sprintf("INSERT INTO %s (shortlink, text, created_at, expiry_in_minutes) VALUES (?, ?, ?, ?)", WRITE_TABLE_NAME)
	_, err := s.db.Exec(query, shortlink, text, createdAt, expiryByMinutes)
	if err != nil {
		return err
	}
	return nil
}

func (s *writePasteRepository) DeleteExpiredContent() error {
	query := fmt.Sprintf("DELETE FROM %s WHERE NOW() > DATE_ADD(created_at, INTERVAL expiry_in_minutes MINUTE) AND expiry_in_minutes != 0", WRITE_TABLE_NAME)
	_, err := s.db.Exec(query)
	return err
}
