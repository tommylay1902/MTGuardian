package dataaccess_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tommylay1902/prescriptionmicro/api/dataaccess"
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
	dao := dataaccess.InitalizePrescriptionService(gormDB)

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
	err := dao.CreatePrescription(prescription)

	// // Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	// Assert that there was no error returned from CreatePrescription
	assert.NoError(t, err)
}

func TestCreatePrescriptionWithDatabaseError(t *testing.T) {
	dao := dataaccess.InitalizePrescriptionService(gormDB)

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

	err := dao.CreatePrescription(prescription2)

	// Check for any errors from the mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	// Assert that there was an error returned from CreatePrescription
	assert.Error(t, err)
}

// Helper functions for creating pointers to string and time values
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}
