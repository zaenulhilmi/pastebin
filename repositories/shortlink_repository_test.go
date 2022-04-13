package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zaenulhilmi/pastebin/entities"
	"github.com/zaenulhilmi/pastebin/repositories"
	"testing"
    "fmt"
)

func Test_FindContentByShortlink(t *testing.T) {
	db, mock, err := sqlmock.New()

    repo := repositories.NewShortlinkRepository(db)
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
    fmt.Println(res)

    assert.NotNil(t, res)
    assert.Nil(t, err)
}


