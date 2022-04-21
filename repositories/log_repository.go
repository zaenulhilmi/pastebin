package repositories

import (
	"database/sql"

	"github.com/zaenulhilmi/pastebin/entities"
)

type LogRepository interface {
	Create(log entities.ShortlinkLog) error
}

func NewLogRepository(db *sql.DB) LogRepository {
	return &logRepository{
		db: db,
	}
}

type logRepository struct {
	db *sql.DB
}

func (l *logRepository) Create(log entities.ShortlinkLog) error {
	_, err := l.db.Exec("INSERT INTO url_visit_histories (url, address, created_at) VALUES (?, ?, ?)", log.Shortlink, log.Address, log.CreatedAt)
	return err
}
