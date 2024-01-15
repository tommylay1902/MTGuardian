package dao_test

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/tommylay1902/authmicro/api/dao"
	"github.com/tommylay1902/authmicro/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db      *sql.DB
	mock    sqlmock.Sqlmock
	gormDB  *gorm.DB
	gormErr error
)

func TestMain(m *testing.M) {
	// Set up the database and SQL mock.
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		panic("Error creating SQL mock: " + err.Error())
	}
	defer db.Close()

	// Create a GORM DB instance from the *gorm.DB
	gormDB, gormErr = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if gormErr != nil {
		panic("Error creating GORM DB: " + gormErr.Error())
	}
	defer func() {
		dbInstance, _ := gormDB.DB()
		_ = dbInstance.Close()
	}()

	// Run the tests and exit
	exitCode := m.Run()

	// Clean up any resources if needed

	// Exit with the status code
	os.Exit(exitCode)
}

func TestRegisterUser(t *testing.T) {

	defer mock.ExpectationsWereMet()
	// Initialize the authdao with the GORM DB
	dao := dao.Initialize(gormDB)

	auth := &model.Auth{
		ID:           uuid.New(),
		Email:        StringPointer("tommylay.c@gmail.com"),
		Password:     StringPointer("$2a$10$/Z8CBBPBv0YlGvfjGglQ3O1mGoftvtF34pXsCmOf6.gvvXYphkO32"),
		RefreshToken: StringPointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbW15bGF5LmNAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMi0xOCAwODoyMzo1NC44MDg3OTE4ICswMDAwIFVUQyIsInN1YiI6InRvbW15bGF5LmNAZ21haWwuY29tIn0.poUKsuF0-ZLV1ky1y-X0h150UAqYZ0MNCYknukfBJDA"),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query and its result in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"auths\" (\"id\",\"email\",\"password\",\"refresh_token\") VALUES ($1,$2,$3,$4)")).
		WithArgs(auth.ID, *auth.Email, *auth.Password, *auth.RefreshToken).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	id, err := dao.CreateAuth(auth)

	// // Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, *id, auth.ID)
}

func TestCreatePrescriptionWithDatabaseError(t *testing.T) {
	defer mock.ExpectationsWereMet()
	dao := dao.Initialize(gormDB)

	auth := &model.Auth{
		ID:           uuid.New(),
		Email:        StringPointer("tommylay.c@gmail.com"),
		Password:     StringPointer("$2a$10$/Z8CBBPBv0YlGvfjGglQ3O1mGoftvtF34pXsCmOf6.gvvXYphkO32"),
		RefreshToken: StringPointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbW15bGF5LmNAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMi0xOCAwODoyMzo1NC44MDg3OTE4ICswMDAwIFVUQyIsInN1YiI6InRvbW15bGF5LmNAZ21haWwuY29tIn0.poUKsuF0-ZLV1ky1y-X0h150UAqYZ0MNCYknukfBJDA"),
	}

	auth2 := &model.Auth{
		ID:           uuid.New(),
		Email:        StringPointer("tommylay.c@gmail.com"),
		Password:     StringPointer("$2a$10$/Z8CBBPBv0YlGvfjGglQ3O1mGoftvtF3kpXsCmOf6.gvvXYphkO32"),
		RefreshToken: StringPointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbW15bGF5LmNAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMi0xOCAwODoyMzo1NC44MDg3OTE4ICswMDAwIFVUQyIsInN1YiI6InavbW15bGF5LmNAZ21haWwuY29tIn0.poUKsuF0-ZLV1ky1y-X0h150UAqYZ0MNCYknukfBJDA"),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query and its result in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"auths\" (\"id\",\"email\",\"password\",\"refresh_token\") VALUES ($1,$2,$3,$4)")).
		WithArgs(auth.ID, *auth.Email, *auth.Password, *auth.RefreshToken).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	dao.CreateAuth(auth)

	mock.ExpectBegin()
	// Set up the expected SQL query and indicate an error in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"auths\" (\"id\",\"email\",\"password\",\"refresh_token\") VALUES ($1,$2,$3,$4)")).
		WithArgs(auth2.ID, *auth2.Email, *auth2.Password, *auth2.RefreshToken).
		WillReturnError(fmt.Errorf("DATABASE WILL THROW ERROR"))
	mock.ExpectRollback()

	_, err := dao.CreateAuth(auth2)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.Error(t, err)
}

func TestGetHashFromEmail(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)
	email := "tommylay.c@gmail.com"
	auth := &model.Auth{
		ID:           uuid.New(),
		Email:        StringPointer(email),
		Password:     StringPointer("$2a$10$/Z8CBBPBv0YlGvfjGglQ3O1mGoftvtF34pXsCmOf6.gvvXYphkO32"),
		RefreshToken: StringPointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbW15bGF5LmNAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMi0xOCAwODoyMzo1NC44MDg3OTE4ICswMDAwIFVUQyIsInN1YiI6InRvbW15bGF5LmNAZ21haWwuY29tIn0.poUKsuF0-ZLV1ky1y-X0h150UAqYZ0MNCYknukfBJDA"),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query and its result in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"auths\" (\"id\",\"email\",\"password\",\"refresh_token\") VALUES ($1,$2,$3,$4)")).
		WithArgs(auth.ID, *auth.Email, *auth.Password, *auth.RefreshToken).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	dao.CreateAuth(auth)

	// Set up the expected SQL query and its result in the mock
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"auths\" WHERE email = $1 ORDER BY \"auths\".\"id\" LIMIT 1")).
		WithArgs(*auth.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "refresh_token"}).
			AddRow(auth.ID, *auth.Email, *auth.Password, *auth.RefreshToken))

	hash, _ := dao.GetHashFromEmail(auth.Email)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.Equal(t, *auth.Password, *hash)

}

func TestGetHashFromEmailWillThrowError(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)
	email := "tommylay.c@gmail.com"

	// Set up the expected SQL query and its result in the mock
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"auths\" WHERE email = $1 ORDER BY \"auths\".\"id\" LIMIT 1")).
		WithArgs(email).WillReturnError(gorm.ErrRecordNotFound)

	hash, err := dao.GetHashFromEmail(&email)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.Empty(t, hash)
	assert.Error(t, err)

}

func TestGetTokenFromEmail(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbW15bGF5LmNAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMi0xOCAwODoyMzo1NC44MDg3OTE4ICswMDAwIFVUQyIsInN1YiI6InRvbW15bGF5LmNAZ21haWwuY29tIn0.poUKsuF0-ZLV1ky1y-X0h150UAqYZ0MNCYknukfBJDA"
	auth := &model.Auth{
		ID:           uuid.New(),
		Email:        StringPointer("tommylay.c@gmail.com"),
		Password:     StringPointer("$2a$10$/Z8CBBPBv0YlGvfjGglQ3O1mGoftvtF34pXsCmOf6.gvvXYphkO32"),
		RefreshToken: StringPointer(token),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query and its result in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"auths\" (\"id\",\"email\",\"password\",\"refresh_token\") VALUES ($1,$2,$3,$4)")).
		WithArgs(auth.ID, *auth.Email, *auth.Password, *auth.RefreshToken).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	dao.CreateAuth(auth)

	// Set up the expected SQL query and its result in the mock
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"auths\" WHERE email = $1 ORDER BY \"auths\".\"id\" LIMIT 1")).
		WithArgs(*auth.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "refresh_token"}).
			AddRow(auth.ID, *auth.Email, *auth.Password, *auth.RefreshToken))

	actual, _ := dao.GetTokenFromEmail(auth.Email)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.Equal(t, *auth.RefreshToken, *actual)

}

func TestGetTokenFromEmailWillThrowError(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)
	email := "tommylay.c@gmail.com"

	// Set up the expected SQL query and its result in the mock
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"auths\" WHERE email = $1 ORDER BY \"auths\".\"id\" LIMIT 1")).
		WithArgs(email).WillReturnError(gorm.ErrRecordNotFound)

	token, err := dao.GetTokenFromEmail(&email)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.Empty(t, token)
	assert.Error(t, err)

}

// Helper functions for creating pointers to string and time values 2
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}
