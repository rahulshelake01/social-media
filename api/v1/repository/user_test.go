package repository

import (
	"database/sql"
	"log"
	"social-media/config"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindUserByEmail(t *testing.T) {
	db, mock := NewMock()
	appConfig := config.GetAppConfig()

	repo := UserRepository(db, appConfig)
	query := "SELECT id, email, pass from users where email = ?"
	rows := sqlmock.NewRows([]string{"id", "email", "pass"}).
		AddRow(10, "test@gm.com", "abc@123")
	mock.ExpectQuery(query).WithArgs("test@gm.com").WillReturnRows(rows)
	user, ifFound, err := repo.FindUserByEmail("test@gm.com")
	assert.NotNil(t, user)
	assert.Equal(t, true, ifFound)
	assert.NoError(t, err)

}
