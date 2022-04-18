package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/mocks"
	"github.com/zaenulhilmi/pastebin/repositories"
)

func TestFindContentByShortlink(t *testing.T) {
	db, mock, err := sqlmock.New()

	clock := helpers.SystemClock{}
	repo := repositories.NewShortlinkRepository(db, clock)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	content := entities.Content{}

	query := "SELECT text, created_at, expiry_in_minutes FROM contents WHERE shortlink = ?"

	rows := sqlmock.NewRows([]string{"text", "created_at", "expiry_in_minutes"}).
		AddRow(&content.Text, &content.CreatedAt, &content.ExpiryInMinutes)

	mock.ExpectQuery(query).
		WithArgs("shortlink").
		WillReturnRows(rows)

	res, err := repo.FindContentByShortlink("shortlink")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestCreateContent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	clock := mocks.ClockMock{}
	repo := repositories.NewShortlinkRepository(db, &clock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT INTO contents (shortlink, text, created_at, expiry_in_minutes) VALUES (?, ?, ?, ?)"
	mock.ExpectExec(query).
		WithArgs("shortlink", "text", clock.Now(), 10).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateContent("shortlink", "text", 10)
	assert.Nil(t, err)
}

func TestDeleteExpiredContent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	clock := mocks.ClockMock{}
	repo := repositories.NewShortlinkRepository(db, &clock)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := "DELETE FROM contents WHERE NOW() > DATE_ADD(created_at, INTERVAL expiry_in_minutes MINUTE) AND expiry_in_minutes != 0"
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteExpiredContent()
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)

}
