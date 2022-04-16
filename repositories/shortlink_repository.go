package repositories

import (
	"database/sql"
	"github.com/zaenulhilmi/pastebin/entities"
)

type ShortlinkRepository interface {
	FindContentByShortlink(shortlink string) (*entities.Content, error)
	CreateContent(shortlink string, text string, expiryByMinutes int) error
}

func NewShortlinkRepository(db *sql.DB) ShortlinkRepository {
	return &shortlinkRepository{
		db: db,
	}
}

type shortlinkRepository struct {
	db *sql.DB
}

func (s *shortlinkRepository) CreateContent(shortlink string, text string, expiryByMinutes int) error {
	//TODO implement me
	panic("implement me")
}

func (s *shortlinkRepository) FindContentByShortlink(shortlink string) (*entities.Content, error) {
	var content entities.Content
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
