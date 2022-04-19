package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/repositories"
)

func TestCreateContent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	clock := mocks.ClockMock{}
	repo := repositories.NewWritePasteRepository(db, &clock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT INTO pastes (shortlink, text, created_at, expiry_in_minutes) VALUES (?, ?, ?, ?)"
	mock.ExpectExec(query).
		WithArgs("shortlink", "text", clock.Now(), 10).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateContent("shortlink", "text", 10)
	assert.Nil(t, err)
}

func TestDeleteExpiredContent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	clock := mocks.ClockMock{}
	repo := repositories.NewWritePasteRepository(db, &clock)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := "DELETE FROM pastes WHERE NOW() > DATE_ADD(created_at, INTERVAL expiry_in_minutes MINUTE) AND expiry_in_minutes != 0"
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteExpiredContent()
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)

}
