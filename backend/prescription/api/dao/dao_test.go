package dao_test

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tommylay1902/prescriptionmicro/api/dao"
	"github.com/tommylay1902/prescriptionmicro/internal/error/apperror"
	"github.com/tommylay1902/prescriptionmicro/internal/model"
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

func TestCreatePrescriptionWithMock(t *testing.T) {

	defer mock.ExpectationsWereMet()
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dao.Initialize(gormDB)

	// Create a sample Prescription
	prescription := &model.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Ended:      TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer("tommylay.c@gmail.com"),
	}
	fmt.Println(*prescription.Started, *prescription.Ended)

	mock.ExpectBegin()
	// Set up the expected SQL query and its result in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"prescriptions\" (\"id\",\"medication\",\"dosage\",\"notes\",\"started\",\"ended\",\"refills\",\"owner\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)")).
		WithArgs(prescription.ID, *prescription.Medication, *prescription.Dosage, *prescription.Notes, *prescription.Started, *prescription.Ended, *prescription.Refills, *prescription.Owner).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	// Call the CreatePrescription method of the DAO
	id, err := dao.CreatePrescription(prescription)

	// // Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	// Assert that there was no error returned from CreatePrescription
	assert.NoError(t, err)
	assert.Equal(t, *id, prescription.ID)
}

func TestCreatePrescriptionWithDatabaseError(t *testing.T) {
	defer mock.ExpectationsWereMet()
	dao := dao.Initialize(gormDB)

	// Create a sample Prescription
	prescription := &model.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now().UTC()),
		Ended:      TimePointer(time.Now().UTC()),
		Refills:    IntPointer(2),
		Owner:      StringPointer("tommylay.c@gmail.com"),
	}

	prescription2 := &model.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now().UTC()),
		Ended:      TimePointer(time.Now().UTC()),
		Refills:    IntPointer(2),
		Owner:      StringPointer("tommylay.c@gmail.com"),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query and indicate an error in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"prescriptions\" (\"id\",\"medication\",\"dosage\",\"notes\",\"started\",\"ended\",\"refills\",\"owner\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)")).WithArgs(
		prescription.ID,
		*prescription.Medication,
		*prescription.Dosage,
		*prescription.Notes,
		*prescription.Started,
		*prescription.Ended,
		*prescription.Refills,
		*prescription.Owner,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// Call the CreatePrescription method of the DAO
	dao.CreatePrescription(prescription)

	mock.ExpectBegin()
	// Set up the expected SQL query and indicate an error in the mock
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"prescriptions\" (\"id\",\"medication\",\"dosage\",\"notes\",\"started\",\"ended\",\"refills\",\"owner\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)")).WithArgs(
		prescription2.ID,
		*prescription2.Medication,
		*prescription2.Dosage,
		*prescription2.Notes,
		*prescription2.Started,
		*prescription2.Ended,
		*prescription2.Refills,
		*prescription2.Owner,
	).WillReturnError(fmt.Errorf("database will throw error"))
	mock.ExpectRollback()

	_, err := dao.CreatePrescription(prescription2)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	// Assert that there was an error returned from CreatePrescription
	assert.Error(t, err)

}

func TestGetPrescriptionById(t *testing.T) {
	defer mock.ExpectationsWereMet()
	dao := dao.Initialize(gormDB)
	id := uuid.New()
	email := "tommylay.c@gmail.com"
	expected := &model.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Ended:      TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer(email),
	}

	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WithArgs(email, id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "medication", "dosage", "notes", "started", "ended", "refills", "owner"}).
			AddRow(expected.ID, *expected.Medication, *expected.Dosage, *expected.Notes, *expected.Started, *expected.Ended, *expected.Refills, *expected.Owner))

	result, err := dao.GetPrescriptionById(id, email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock %v", err)
	}
	assert.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestGetPrescriptionByIdInvalidId(t *testing.T) {
	defer mock.ExpectationsWereMet()
	dao := dao.Initialize(gormDB)
	id := uuid.New()
	email := "tommylay.c@gmail.com"

	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WithArgs(email, id).WillReturnError(gorm.ErrRecordNotFound)

	result, err := dao.GetPrescriptionById(id, email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock %v", err)
	}

	assert.Nil(t, result)

	assert.Error(t, err)
	// expectedErr := &apperror.ResourceNotFound{Code: 404}

	assert.True(t, errors.Is(err, &apperror.ResourceNotFound{Code: 404}))
}

func TestGetAllPrescriptions(t *testing.T) {
	defer mock.ExpectationsWereMet()
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dao.Initialize(gormDB)
	email := "tommylay.c@gmail.com"

	// Create two sample Prescription instances
	id1 := uuid.New()
	expectedOne := &model.Prescription{
		ID:         id1,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer(email),
	}

	id2 := uuid.New()
	expectedTwo := &model.Prescription{
		ID:         id2,
		Medication: StringPointer("Sample Medication 2"),
		Dosage:     StringPointer("Sample Dosage 2"),
		Notes:      StringPointer("Sample Notes 2"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer(email),
	}

	// Set up the expected SQL query and its result in the mock to return both expectedOne and expectedTwo
	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WillReturnRows(sqlmock.NewRows([]string{"id", "medication", "dosage", "notes", "started", "refills", "owner"}).
			AddRow(expectedTwo.ID, *expectedTwo.Medication, *expectedTwo.Dosage, *expectedTwo.Notes, *expectedTwo.Started, expectedTwo.Refills, *expectedTwo.Owner).
			AddRow(expectedOne.ID, *expectedOne.Medication, *expectedOne.Dosage, *expectedOne.Notes, *expectedOne.Started, *expectedOne.Refills, *expectedOne.Owner),
		)

	// Call the GetAllPrescriptions method of the DAO
	results, err := dao.GetAllPrescriptions(make(map[string]string), &email)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	// Assert that there was no error returned from GetAllPrescriptions
	assert.NoError(t, err)

	// Assert that the results contain both expectedOne and expectedTwo
	assert.Contains(t, results, *expectedOne)
	assert.Contains(t, results, *expectedTwo)
}

func TestGetAllPrescriptionsWithError(t *testing.T) {
	defer mock.ExpectationsWereMet()
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dao.Initialize(gormDB)
	email := "emailDNE@gmail.com"

	// Set up the mock to expect an error when querying for prescriptions
	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WithArgs().
		WillReturnError(fmt.Errorf("database error"))

	// Call the GetAllPrescriptions method of the DAO
	_, err := dao.GetAllPrescriptions(make(map[string]string), &email)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	// Assert that there was an error returned from GetAllPrescriptions
	assert.Error(t, err)
}

func TestDeletePrescription(t *testing.T) {
	defer mock.ExpectationsWereMet()
	dao := dao.Initialize(gormDB)
	id := uuid.New()
	email := "tommylay.c@gmail.com"

	expected := &model.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer(email),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query for DELETE using ExpectExec
	mock.ExpectExec("DELETE FROM \"prescriptions\"").
		WithArgs(email, id).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Indicate that one row was affected
	mock.ExpectCommit()
	// Call the DeletePrescription method of the DAO
	err := dao.DeletePrescription(expected, email)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	// Assert that there was no error returned from DeletePrescription
	assert.NoError(t, err)
}

func TestDeletePrescriptionWithError(t *testing.T) {
	defer mock.ExpectationsWereMet()
	dao := dao.Initialize(gormDB)
	id := uuid.New()
	email := "tommylay.c@gmail.com"

	expected := &model.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Owner:      StringPointer(email),
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"prescriptions\"").
		WithArgs(email, id).
		WillReturnError(fmt.Errorf("Database Error"))
	mock.ExpectRollback()
	err := dao.DeletePrescription(expected, email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	assert.Error(t, err)
}

func TestUpdatePrescription(t *testing.T) {
	defer mock.ExpectationsWereMet()
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dao.Initialize(gormDB)

	// Create a sample Prescription
	prescription := &model.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Ended:      TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer("tommylay.c@gmail.com"),
	}

	// Insert the initial prescription into the database
	dao.CreatePrescription(prescription)

	// Update the prescription with new values
	updatedPrescription := &model.Prescription{
		ID:         prescription.ID,
		Medication: StringPointer("Updated Medication"),
		Dosage:     StringPointer("Updated Dosage"),
		Notes:      StringPointer("Updated Notes"),
		Started:    TimePointer(time.Now()),
		Ended:      TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      prescription.Owner,
	}

	mock.ExpectBegin()
	// Set up the expected SQL query for the update operation
	mock.ExpectExec(`UPDATE "prescriptions"`).
		WithArgs(*updatedPrescription.Medication, *updatedPrescription.Dosage, *updatedPrescription.Notes, *updatedPrescription.Started, *updatedPrescription.Ended, *prescription.Refills, *prescription.Owner, *prescription.Owner, prescription.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Indicate that one row was updated
	mock.ExpectCommit()

	err := dao.UpdatePrescription(updatedPrescription, *prescription.Owner)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	// Assert that there was no error returned from UpdatePrescription
	assert.NoError(t, err)

}

func TestUpdatePrescriptionWithError(t *testing.T) {
	defer mock.ExpectationsWereMet()
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dao.Initialize(gormDB)

	// Create a sample Prescription
	prescription := &model.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Ended:      nil,
		Refills:    IntPointer(2),
		Owner:      StringPointer("tommylay.c@gmail.com"),
	}

	// Insert the initial prescription into the database
	dao.CreatePrescription(prescription)

	// Update the prescription with new values
	updatedPrescription := &model.Prescription{
		ID:         prescription.ID,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    prescription.Started,
		Ended:      TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      prescription.Owner,
	}

	mock.ExpectBegin()
	// Set up the expected SQL query for the update operation
	mock.ExpectExec("UPDATE \"prescriptions\"").
		WithArgs(*updatedPrescription.Medication, *updatedPrescription.Dosage, *updatedPrescription.Notes, *updatedPrescription.Started, *updatedPrescription.Ended, *prescription.Refills, *prescription.Owner, *prescription.Owner, prescription.ID).
		WillReturnError(fmt.Errorf("database error"))
	mock.ExpectRollback()

	err := dao.UpdatePrescription(updatedPrescription, *prescription.Owner)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	assert.Error(t, err)
}

// Helper functions for creating pointers to string and time values 2
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}

func IntPointer(i int) *int {
	return &i
}
