package dataaccess_test

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/tommylay1902/authmicro/api/dataaccess"
	"github.com/tommylay1902/authmicro/internal/models"

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
	dao := dataaccess.InitializeAuthDAO(gormDB)

	auth := &models.Auth{
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
	dao := dataaccess.InitializeAuthDAO(gormDB)

	auth := &models.Auth{
		ID:           uuid.New(),
		Email:        StringPointer("tommylay.c@gmail.com"),
		Password:     StringPointer("$2a$10$/Z8CBBPBv0YlGvfjGglQ3O1mGoftvtF34pXsCmOf6.gvvXYphkO32"),
		RefreshToken: StringPointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbW15bGF5LmNAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMi0xOCAwODoyMzo1NC44MDg3OTE4ICswMDAwIFVUQyIsInN1YiI6InRvbW15bGF5LmNAZ21haWwuY29tIn0.poUKsuF0-ZLV1ky1y-X0h150UAqYZ0MNCYknukfBJDA"),
	}

	auth2 := &models.Auth{
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

	dao := dataaccess.InitializeAuthDAO(gormDB)
	email := "tommylay.c@gmail.com"
	auth := &models.Auth{
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

	dao := dataaccess.InitializeAuthDAO(gormDB)
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

	dao := dataaccess.InitializeAuthDAO(gormDB)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbW15bGF5LmNAZ21haWwuY29tIiwiZXhwIjoiMjAyMy0xMi0xOCAwODoyMzo1NC44MDg3OTE4ICswMDAwIFVUQyIsInN1YiI6InRvbW15bGF5LmNAZ21haWwuY29tIn0.poUKsuF0-ZLV1ky1y-X0h150UAqYZ0MNCYknukfBJDA"
	auth := &models.Auth{
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

	dao := dataaccess.InitializeAuthDAO(gormDB)
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

// func TestGetAllPrescriptions(t *testing.T) {
// 	defer mock.ExpectationsWereMet()
// 	// Initialize the PrescriptionDAO with the GORM DB
// 	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

// 	// Create two sample Prescription instances
// 	id1 := uuid.New()
// 	expectedOne := &models.Prescription{
// 		ID:         id1,
// 		Medication: StringPointer("Sample Medication"),
// 		Dosage:     StringPointer("Sample Dosage"),
// 		Notes:      StringPointer("Sample Notes"),
// 		Started:    TimePointer(time.Now()),
// 	}

// 	id2 := uuid.New()
// 	expectedTwo := &models.Prescription{
// 		ID:         id2,
// 		Medication: StringPointer("Sample Medication 2"),
// 		Dosage:     StringPointer("Sample Dosage 2"),
// 		Notes:      StringPointer("Sample Notes 2"),
// 		Started:    TimePointer(time.Now()),
// 	}

// 	// Set up the expected SQL query and its result in the mock to return both expectedOne and expectedTwo
// 	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "medication", "dosage", "notes", "started"}).
// 			AddRow(expectedTwo.ID, *expectedTwo.Medication, *expectedTwo.Dosage, *expectedTwo.Notes, *expectedTwo.Started).
// 			AddRow(expectedOne.ID, *expectedOne.Medication, *expectedOne.Dosage, *expectedOne.Notes, *expectedOne.Started),
// 		)

// 	// Call the GetAllPrescriptions method of the DAO
// 	results, err := dao.GetAllPrescriptions(make(map[string]string))

// 	// Check for any errors from the mock expectations
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Error in SQL mock: %v", err)
// 	}

// 	// Assert that there was no error returned from GetAllPrescriptions
// 	assert.NoError(t, err)

// 	// Assert that the results contain both expectedOne and expectedTwo
// 	assert.Contains(t, results, *expectedOne)
// 	assert.Contains(t, results, *expectedTwo)
// }

// func TestGetAllPrescriptionsWithError(t *testing.T) {
// 	defer mock.ExpectationsWereMet()
// 	// Initialize the PrescriptionDAO with the GORM DB
// 	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

// 	// Set up the mock to expect an error when querying for prescriptions
// 	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
// 		WithArgs().
// 		WillReturnError(fmt.Errorf("database error"))

// 	// Call the GetAllPrescriptions method of the DAO
// 	_, err := dao.GetAllPrescriptions(make(map[string]string))

// 	// Check for any errors from the mock expectations
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Error in SQL mock: %v", err)
// 	}

// 	// Assert that there was an error returned from GetAllPrescriptions
// 	assert.Error(t, err)
// }

// func TestDeletePrescription(t *testing.T) {
// 	defer mock.ExpectationsWereMet()
// 	dao := dataaccess.InitalizePrescriptionDAO(gormDB)
// 	id := uuid.New()

// 	expected := &models.Prescription{
// 		ID:         id,
// 		Medication: StringPointer("Sample Medication"),
// 		Dosage:     StringPointer("Sample Dosage"),
// 		Notes:      StringPointer("Sample Notes"),
// 		Started:    TimePointer(time.Now()),
// 	}

// 	mock.ExpectBegin()
// 	// Set up the expected SQL query for DELETE using ExpectExec
// 	mock.ExpectExec("DELETE FROM \"prescriptions\" WHERE \"prescriptions\".\"id\" = ?").
// 		WithArgs(id).
// 		WillReturnResult(sqlmock.NewResult(0, 1)) // Indicate that one row was affected
// 	mock.ExpectCommit()
// 	// Call the DeletePrescription method of the DAO
// 	err := dao.DeletePrescription(expected)

// 	// Check for any errors from the mock expectations
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Error in SQL mock: %v", err)
// 	}

// 	// Assert that there was no error returned from DeletePrescription
// 	assert.NoError(t, err)
// }

// func TestDeletePrescriptionWithError(t *testing.T) {
// 	defer mock.ExpectationsWereMet()
// 	dao := dataaccess.InitalizePrescriptionDAO(gormDB)
// 	id := uuid.New()

// 	expected := &models.Prescription{
// 		ID:         id,
// 		Medication: StringPointer("Sample Medication"),
// 		Dosage:     StringPointer("Sample Dosage"),
// 		Notes:      StringPointer("Sample Notes"),
// 		Started:    TimePointer(time.Now()),
// 	}

// 	mock.ExpectBegin()
// 	mock.ExpectExec("DELETE FROM \"prescriptions\" WHERE \"prescriptions\".\"id\" = ?").
// 		WithArgs(id).
// 		WillReturnError(fmt.Errorf("Database Error"))
// 	mock.ExpectRollback()
// 	err := dao.DeletePrescription(expected)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Error in SQL mock: %v", err)
// 	}

// 	assert.Error(t, err)
// }

// func TestUpdatePrescription(t *testing.T) {
// 	defer mock.ExpectationsWereMet()
// 	// Initialize the PrescriptionDAO with the GORM DB
// 	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

// 	// Create a sample Prescription
// 	prescription := &models.Prescription{
// 		ID:         uuid.New(),
// 		Medication: StringPointer("Sample Medication"),
// 		Dosage:     StringPointer("Sample Dosage"),
// 		Notes:      StringPointer("Sample Notes"),
// 		Started:    TimePointer(time.Now()),
// 		Ended:      TimePointer(time.Now()),
// 	}

// 	// Insert the initial prescription into the database
// 	dao.CreatePrescription(prescription)

// 	// Update the prescription with new values
// 	updatedPrescription := &models.Prescription{
// 		ID:         prescription.ID,
// 		Medication: StringPointer("Updated Medication"),
// 		Dosage:     StringPointer("Updated Dosage"),
// 		Notes:      StringPointer("Updated Notes"),
// 		Started:    TimePointer(time.Now()),
// 		Ended:      TimePointer(time.Now()),
// 	}

// 	mock.ExpectBegin()
// 	// Set up the expected SQL query for the update operation
// 	mock.ExpectExec("UPDATE \"prescriptions\" SET .* WHERE \"id\" = ?").
// 		WithArgs(*updatedPrescription.Medication, *updatedPrescription.Dosage, *updatedPrescription.Notes, *updatedPrescription.Started, *updatedPrescription.Ended, prescription.ID).
// 		WillReturnResult(sqlmock.NewResult(0, 1)) // Indicate that one row was updated
// 	mock.ExpectCommit()

// 	err := dao.UpdatePrescription(updatedPrescription)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Error in SQL mock: %v", err)
// 	}

// 	// Assert that there was no error returned from UpdatePrescription
// 	assert.NoError(t, err)

// }

// func TestUpdatePrescriptionWithError(t *testing.T) {
// 	defer mock.ExpectationsWereMet()
// 	// Initialize the PrescriptionDAO with the GORM DB
// 	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

// 	// Create a sample Prescription
// 	prescription := &models.Prescription{
// 		ID:         uuid.New(),
// 		Medication: StringPointer("Sample Medication"),
// 		Dosage:     StringPointer("Sample Dosage"),
// 		Notes:      StringPointer("Sample Notes"),
// 		Started:    TimePointer(time.Now()),
// 		Ended:      nil,
// 	}

// 	// Insert the initial prescription into the database
// 	dao.CreatePrescription(prescription)

// 	// Update the prescription with new values
// 	updatedPrescription := &models.Prescription{
// 		ID:         prescription.ID,
// 		Medication: StringPointer("Sample Medication"),
// 		Dosage:     StringPointer("Sample Dosage"),
// 		Notes:      StringPointer("Sample Notes"),
// 		Started:    prescription.Started,
// 		Ended:      TimePointer(time.Now()),
// 	}

// 	mock.ExpectBegin()
// 	// Set up the expected SQL query for the update operation
// 	mock.ExpectExec("UPDATE \"prescriptions\" SET .* WHERE \"id\" = ?").
// 		WithArgs(*updatedPrescription.Medication, *updatedPrescription.Dosage, *updatedPrescription.Notes, *updatedPrescription.Started, *updatedPrescription.Ended, prescription.ID).
// 		WillReturnError(fmt.Errorf("database error"))
// 	mock.ExpectRollback()

// 	err := dao.UpdatePrescription(updatedPrescription)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Error in SQL mock: %v", err)
// 	}

// 	assert.Error(t, err)
// }

// Helper functions for creating pointers to string and time values 2
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}
