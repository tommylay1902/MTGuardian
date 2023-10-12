package dataaccess_test

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tommylay1902/prescriptionmicro/api/dataaccess"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
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

	// Initialize the PrescriptionDAO with the GORM DB
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

	// Create a sample Prescription
	prescription := &models.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query and its result in the mock
	mock.ExpectExec("INSERT INTO \"prescriptions\"").WithArgs(
		prescription.ID,
		*prescription.Medication,
		*prescription.Dosage,
		*prescription.Notes,
		*prescription.Started,
	).WillReturnResult(sqlmock.NewResult(1, 1))

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
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

	// Create a sample Prescription
	prescription := &models.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	prescription2 := &models.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query and indicate an error in the mock
	mock.ExpectExec("INSERT INTO \"prescriptions\"").WithArgs(
		prescription.ID,
		*prescription.Medication,
		*prescription.Dosage,
		*prescription.Notes,
		*prescription.Started,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// Call the CreatePrescription method of the DAO
	dao.CreatePrescription(prescription)

	mock.ExpectBegin()
	// Set up the expected SQL query and indicate an error in the mock
	mock.ExpectExec("INSERT INTO \"prescriptions\"").WithArgs(
		prescription2.ID,
		*prescription2.Medication,
		*prescription2.Dosage,
		*prescription2.Notes,
		*prescription2.Started,
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
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)
	id := uuid.New()
	expected := &models.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "medication", "dosage", "notes", "started"}).
			AddRow(expected.ID, *expected.Medication, *expected.Dosage, *expected.Notes, *expected.Started))

	result, err := dao.GetPrescriptionById(id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock %v", err)
	}
	assert.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestGetPrescriptionByIdInvalidId(t *testing.T) {
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)
	id := uuid.New()
	// expected := &models.Prescription{
	// 	ID:         id,
	// 	Medication: StringPointer("Sample Medication"),
	// 	Dosage:     StringPointer("Sample Dosage"),
	// 	Notes:      StringPointer("Sample Notes"),
	// 	Started:    TimePointer(time.Now()),
	// }

	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WithArgs(id).WillReturnError(gorm.ErrRecordNotFound)

	result, err := dao.GetPrescriptionById(id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock %v", err)
	}

	assert.Nil(t, result)

	assert.Error(t, err)
	// expectedErr := &customerrors.ResourceNotFound{Code: 404}

	assert.True(t, errors.Is(err, &customerrors.ResourceNotFound{Code: 404}))
}

func TestGetAllPrescriptions(t *testing.T) {
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

	// Create two sample Prescription instances
	id1 := uuid.New()
	expectedOne := &models.Prescription{
		ID:         id1,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	id2 := uuid.New()
	expectedTwo := &models.Prescription{
		ID:         id2,
		Medication: StringPointer("Sample Medication 2"),
		Dosage:     StringPointer("Sample Dosage 2"),
		Notes:      StringPointer("Sample Notes 2"),
		Started:    TimePointer(time.Now()),
	}

	// Set up the expected SQL query and its result in the mock to return both expectedOne and expectedTwo
	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WillReturnRows(sqlmock.NewRows([]string{"id", "medication", "dosage", "notes", "started"}).
			AddRow(expectedTwo.ID, *expectedTwo.Medication, *expectedTwo.Dosage, *expectedTwo.Notes, *expectedTwo.Started).
			AddRow(expectedOne.ID, *expectedOne.Medication, *expectedOne.Dosage, *expectedOne.Notes, *expectedOne.Started),
		)

	// Call the GetAllPrescriptions method of the DAO
	results, err := dao.GetAllPrescriptions()

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
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

	// Set up the mock to expect an error when querying for prescriptions
	mock.ExpectQuery("SELECT .* FROM \"prescriptions\"").
		WillReturnError(fmt.Errorf("database error"))

	// Call the GetAllPrescriptions method of the DAO
	_, err := dao.GetAllPrescriptions()

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	// Assert that there was an error returned from GetAllPrescriptions
	assert.Error(t, err)
}

func TestDeletePrescription(t *testing.T) {
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)
	id := uuid.New()

	expected := &models.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query for DELETE using ExpectExec
	mock.ExpectExec("DELETE FROM \"prescriptions\" WHERE \"prescriptions\".\"id\" = ?").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Indicate that one row was affected
	mock.ExpectCommit()
	// Call the DeletePrescription method of the DAO
	err := dao.DeletePrescription(expected)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	// Assert that there was no error returned from DeletePrescription
	assert.NoError(t, err)
}

func TestDeletePrescriptionWithError(t *testing.T) {
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)
	id := uuid.New()

	expected := &models.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"prescriptions\" WHERE \"prescriptions\".\"id\" = ?").
		WithArgs(id).
		WillReturnError(fmt.Errorf("Database Error"))
	mock.ExpectRollback()
	err := dao.DeletePrescription(expected)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	assert.Error(t, err)
}

func TestUpdatePrescription(t *testing.T) {
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

	// Create a sample Prescription
	prescription := &models.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	// Insert the initial prescription into the database
	dao.CreatePrescription(prescription)

	// Update the prescription with new values
	updatedPrescription := &models.Prescription{
		ID:         prescription.ID,
		Medication: StringPointer("Updated Medication"),
		Dosage:     StringPointer("Updated Dosage"),
		Notes:      StringPointer("Updated Notes"),
		Started:    TimePointer(time.Now()),
	}

	mock.ExpectBegin()
	// Set up the expected SQL query for the update operation
	mock.ExpectExec("UPDATE \"prescriptions\" SET .* WHERE \"id\" = ?").
		WithArgs(*updatedPrescription.Medication, *updatedPrescription.Dosage, *updatedPrescription.Notes, *updatedPrescription.Started, prescription.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Indicate that one row was updated
	mock.ExpectCommit()

	err := dao.UpdatePrescription(updatedPrescription)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	// Assert that there was no error returned from UpdatePrescription
	assert.NoError(t, err)

}

func TestUpdatePrescriptionWithError(t *testing.T) {
	// Initialize the PrescriptionDAO with the GORM DB
	dao := dataaccess.InitalizePrescriptionDAO(gormDB)

	// Create a sample Prescription
	prescription := &models.Prescription{
		ID:         uuid.New(),
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	// Insert the initial prescription into the database
	dao.CreatePrescription(prescription)

	// Update the prescription with new values
	updatedPrescription := &models.Prescription{
		ID:         prescription.ID,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    prescription.Started,
	}

	mock.ExpectBegin()
	// Set up the expected SQL query for the update operation
	mock.ExpectExec("UPDATE \"prescriptions\" SET .* WHERE \"id\" = ?").
		WithArgs(*updatedPrescription.Medication, *updatedPrescription.Dosage, *updatedPrescription.Notes, *updatedPrescription.Started, prescription.ID).
		WillReturnError(fmt.Errorf("data base error"))
	mock.ExpectRollback()

	err := dao.UpdatePrescription(updatedPrescription)

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
