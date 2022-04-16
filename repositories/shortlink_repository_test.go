package repositories_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/repositories"
	"testing"
	"time"
)

func Test_FindContentByShortlink(t *testing.T) {
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

type mockClock struct{}

func (m *mockClock) Now() time.Time {
	return time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func TestCreateContent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	clock := mockClock{}
	repo := repositories.NewShortlinkRepository(db, &clock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// expect execute query
	query := "INSERT INTO contents (shortlink, text, created_at, expiry_in_minutes) VALUES (?, ?, ?, ?)"
	mock.ExpectExec(query).
		WithArgs("shortlink", "text", clock.Now(), 10).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateContent("shortlink", "text", 10)
	assert.Nil(t, err)
}
