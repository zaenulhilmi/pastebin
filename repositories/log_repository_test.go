package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/repositories"
)

func TestLogCreate(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	repo := repositories.NewLogRepository(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	log := entities.ShortlinkLog{
		Shortlink: "http://google.com",
		Address:   "abcd",
	}

	query := "INSERT INTO url_visit_histories (url, address, created_at) VALUES (?, ?, ?)"
	mock.ExpectExec(query).
		WithArgs(log.Shortlink, log.Address, log.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(log)
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}
