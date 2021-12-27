package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dyleme/image-coverter/internal/model"
	"github.com/Dyleme/image-coverter/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		testName string
		user     model.User
		repoErr  error
		wantID   int
		wantErr  error
	}{
		{
			testName: "all is good",
			user: model.User{
				Nickname: "alekse",
				Email:    "my@email.com",
				Password: "password",
			},
			repoErr: nil,
			wantID:  12,
			wantErr: nil,
		},
	}

	query := fmt.Sprintf(`INSERT INTO %s (nickname, email, password_hash)
	 VALUES ($1, $2, $3) RETURNING id`, repository.UsersTable)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			repo := repository.NewAuthPostgres(db)

			rows := sqlmock.NewRows([]string{"id"}).AddRow(tc.wantID)

			mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(tc.user.Nickname, tc.user.Email, tc.user.Password).
				WillReturnRows(rows)

			gotID, gotErr := repo.CreateUser(context.Background(), tc.user)

			assert.ErrorIs(t, gotErr, tc.wantErr)
			assert.Equal(t, gotID, tc.wantID)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were fulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetPasswordAndID(t *testing.T) {
	testCases := []struct {
		testName     string
		userID       int
		userNickname string
		userPassword string
		wantID       int
		wantPassword []byte
		wantErr      error
	}{
		{
			testName:     "all is good",
			userID:       12,
			userNickname: "alekse",
			userPassword: "password",
			wantID:       12,
			wantPassword: []byte("password"),
			wantErr:      nil,
		},
		{
			testName:     "unknown nickname",
			userID:       12,
			userNickname: "unknown",
			userPassword: "password",
			wantID:       0,
			wantPassword: nil,
			wantErr:      sql.ErrNoRows,
		},
	}

	query := fmt.Sprintf("SELECT password_hash, id FROM %s WHERE nickname = ?", repository.UsersTable)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			repo := repository.NewAuthPostgres(db)

			rows := sqlmock.NewRows([]string{"password_hash", "id"})
			if tc.wantPassword != nil && tc.wantID != 0 {
				rows = rows.AddRow(tc.userPassword, tc.userID)
			}
			mock.ExpectQuery(query).WithArgs(tc.userNickname).WillReturnRows(rows)

			gotPassword, gotID, gotErr := repo.GetPasswordHashAndID(context.Background(), tc.userNickname)

			assert.Equal(t, gotID, tc.wantID)
			assert.Equal(t, gotPassword, tc.wantPassword)
			assert.ErrorIs(t, gotErr, tc.wantErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were fulfilled expectations: %s", err)
			}
		})
	}
}
