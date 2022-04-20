package repositories

import (
	"database/sql"

	"github.com/zaenulhilmi/pastebin/entities"
)

type LogRepository interface {
	Create(log entities.Log) error
}

func NewLogRepository(db *sql.DB) LogRepository {
	return &logRepository{
		db: db,
	}
}

type logRepository struct {
	db *sql.DB
}

func (l *logRepository) Create(log entities.Log) error {
	_, err := l.db.Exec("INSERT INTO url_visit_histories (url, address, method, created_at) VALUES (?, ?, ?, ?)", log.Url, log.Address, log.Method, log.CreatedAt)
	return err
}
