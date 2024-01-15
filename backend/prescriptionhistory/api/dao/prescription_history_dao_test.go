package dao_test

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tommylay1902/prescriptionhistory/api/dao"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
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

func TestCreateHistory(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)

	id := uuid.New()
	rxId := uuid.New()
	owner := "tommylay.c@gmail.com"
	taken := time.Now()

	rxHistory := &model.PrescriptionHistory{
		Id:             id,
		PrescriptionId: rxId,
		Owner:          owner,
		Taken:          &taken,
	}

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"prescription_histories\" (\"id\",\"prescription_id\",\"owner\",\"taken\") VALUES ($1,$2,$3,$4)")).WithArgs(rxHistory.Id, rxHistory.PrescriptionId, rxHistory.Owner, rxHistory.Taken).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	result, err := dao.CreateHistory(rxHistory)

	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.Equal(t, *result, rxHistory.Id)

}

func TestCreateHistoryWithDatabaseError(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)

	id := uuid.New()
	rxId := uuid.New()
	owner := "tommylay.c@gmail.com"
	taken := time.Now()

	rxHistory := &model.PrescriptionHistory{
		Id:             id,
		PrescriptionId: rxId,
		Owner:          owner,
		Taken:          &taken,
	}

	rxHistoryDup := &model.PrescriptionHistory{
		Id:             id,
		PrescriptionId: rxId,
		Owner:          owner,
		Taken:          &taken,
	}

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"prescription_histories\" (\"id\",\"prescription_id\",\"owner\",\"taken\") VALUES ($1,$2,$3,$4)")).WithArgs(rxHistory.Id, rxHistory.PrescriptionId, rxHistory.Owner, rxHistory.Taken).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	result, err := dao.CreateHistory(rxHistory)

	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	assert.Equal(t, *result, rxHistory.Id)

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"prescription_histories\" (\"id\",\"prescription_id\",\"owner\",\"taken\") VALUES ($1,$2,$3,$4)")).WithArgs(rxHistoryDup.Id, rxHistoryDup.PrescriptionId, rxHistoryDup.Owner, rxHistoryDup.Taken).WillReturnError(fmt.Errorf("database will throw error"))

	mock.ExpectRollback()

	_, expectErr := dao.CreateHistory(rxHistoryDup)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock expectations: %v", err)
	}

	// Assert that there was an error returned from CreatePrescription
	assert.Error(t, expectErr)
}

func TestGetPrescriptionHistory(t *testing.T) {
	defer mock.ExpectationsWereMet()
	id := uuid.New()
	rxId := uuid.New()
	owner := "tommylay.c@gmail.com"
	taken := time.Now()

	rxHistory := &model.PrescriptionHistory{
		Id:             id,
		PrescriptionId: rxId,
		Owner:          owner,
		Taken:          &taken,
	}

	id2 := uuid.New()
	rxId2 := uuid.New()
	owner2 := "tommylay.d@gmail.com"
	taken2 := time.Now()

	rxHistory2 := &model.PrescriptionHistory{
		Id:             id2,
		PrescriptionId: rxId2,
		Owner:          owner2,
		Taken:          &taken2,
	}

	dao := dao.Initialize(gormDB)

	mock.ExpectQuery("SELECT .* FROM \"prescription_histories\"").WillReturnRows(sqlmock.NewRows([]string{"id", "prescription_id", "owner", "taken"}).AddRow(
		rxHistory.Id, rxHistory.PrescriptionId, rxHistory.Owner, rxHistory.Taken,
	).AddRow(
		rxHistory2.Id, rxHistory2.PrescriptionId, rxHistory2.Owner, rxHistory2.Taken,
	))

	results, err := dao.GetPrescriptionHistory(make(map[string]string), owner)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

	assert.NoError(t, err)

	assert.Contains(t, results, *rxHistory)
	assert.Contains(t, results, *rxHistory2)
}

func TestGetPrescriptionHistoryWithError(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)

	mock.ExpectQuery("SELECT .* FROM \"prescription_histories\"").WithArgs("tommylay.d@gmail.com").WillReturnError(gorm.ErrRecordNotFound)

	_, err := dao.GetPrescriptionHistory(make(map[string]string), "tommylay.d@gmail.com")

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}
}

func TestGetByEmailAndRx(t *testing.T) {
	defer mock.ExpectationsWereMet()
	taken := time.Now()
	email := "tommylay.c@gmail.com"
	id := uuid.New()

	expected := &model.PrescriptionHistory{
		Id:             uuid.New(),
		PrescriptionId: id,
		Owner:          email,
		Taken:          &taken,
	}

	dao := dao.Initialize(gormDB)

	mock.ExpectQuery("SELECT .* FROM \"prescription_histories\"").WithArgs(email, id).WillReturnRows(sqlmock.NewRows([]string{"id", "prescription_id", "owner", "taken"}).AddRow(expected.Id, expected.PrescriptionId, expected.Owner, expected.Taken))

	result, err := dao.GetByEmailAndRx(email, id)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock %v", err)
	}

	assert.Equal(t, *expected, *result)
}

func TestGetByInvalidEmailAndInvalidRx(t *testing.T) {
	defer mock.ExpectationsWereMet()
	dao := dao.Initialize(gormDB)

	mock.ExpectQuery("SELECT .* FROM \"prescription_histories\"").WillReturnError(gorm.ErrRecordNotFound)

	result, err := dao.GetByEmailAndRx("", uuid.New())

	assert.Nil(t, result)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock %v", err)
	}
}

func TestDeleteByIdWithError(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)
	id := uuid.New()
	email := "tommylay.c@gmail.com"

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"prescription_histories\"").WithArgs(email, id).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	err := dao.DeleteByEmailAndId(email, id)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}
}

func TestDeleteByInvalidIdAndEmail(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)
	id := uuid.New()
	email := "tommylay.c@gmail.com"

	mock.ExpectBegin()

	mock.ExpectExec("DELETE FROM \"prescription_histories\"").WithArgs(email, id).WillReturnError(fmt.Errorf("Database error"))

	mock.ExpectRollback()

	err := dao.DeleteByEmailAndId(email, id)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)

	taken := time.Now()
	email := "tommylay.c@gmail.com"

	rxHistory := &model.PrescriptionHistory{
		Id:             uuid.New(),
		PrescriptionId: uuid.New(),
		Owner:          email,
		Taken:          &taken,
	}

	dao.CreateHistory(rxHistory)

	takenUpdate := time.Now()
	updatedRx := &model.PrescriptionHistory{
		Id:             rxHistory.Id,
		PrescriptionId: rxHistory.PrescriptionId,
		Owner:          rxHistory.Owner,
		Taken:          &takenUpdate,
	}

	mock.ExpectBegin()

	mock.ExpectExec(`UPDATE "prescription_histories"`).WithArgs(updatedRx.Id, updatedRx.PrescriptionId, updatedRx.Owner, *updatedRx.Taken, updatedRx.Owner, updatedRx.PrescriptionId, updatedRx.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	err := dao.UpdateByEmailAndRx(*updatedRx, email, rxHistory.PrescriptionId)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}

}

func TestUpdateNoChange(t *testing.T) {
	defer mock.ExpectationsWereMet()

	dao := dao.Initialize(gormDB)
	taken := time.Now()
	email := "tommylay.c@gmail.com"

	rxHistory := &model.PrescriptionHistory{
		Id:             uuid.New(),
		PrescriptionId: uuid.New(),
		Owner:          email,
		Taken:          &taken,
	}

	dao.CreateHistory(rxHistory)

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE \"prescription_histories\"").WithArgs(rxHistory.Id, rxHistory.PrescriptionId, rxHistory.Owner, *rxHistory.Taken, rxHistory.Owner, rxHistory.PrescriptionId, rxHistory.Id).WillReturnError(fmt.Errorf("database error"))

	mock.ExpectRollback()

	err := dao.UpdateByEmailAndRx(*rxHistory, rxHistory.Owner, rxHistory.PrescriptionId)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Error in SQL mock: %v", err)
	}
}
