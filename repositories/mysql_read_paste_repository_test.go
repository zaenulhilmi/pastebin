package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/repositories"
)

func TestFindContentByShortlink(t *testing.T) {
	db, mock, err := sqlmock.New()

	clock := helpers.SystemClock{}
	repo := repositories.NewReadPasteRepository(db, clock)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	content := entities.Paste{}

	query := "SELECT text, created_at, expiry_in_minutes FROM pastes WHERE shortlink = ?"

	rows := sqlmock.NewRows([]string{"text", "created_at", "expiry_in_minutes"}).
		AddRow(&content.Text, &content.CreatedAt, &content.ExpiryInMinutes)

	mock.ExpectQuery(query).
		WithArgs("shortlink").
		WillReturnRows(rows)

	res, err := repo.FindContentByShortlink("shortlink")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
