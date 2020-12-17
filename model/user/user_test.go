package user

import (
	"database/sql"
	"testing"
)

type dbMock struct {
	db       *sql.DB
	mockFunc func() *sql.DB
}

func (sql *dbMock) New()             {}
func (sql *dbMock) GetDB() *sql.DB   { return sql.mockFunc() }
func (sql *dbMock) SetDB(sq *sql.DB) { sql.db = sq }

// @TODO
// done model tests
func TestRegister(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()

	// dm := dbMock{}
	// dm.SetDB(db)
	// postgres.DB = &dm

	// Model.Set(&User{FirstName: "name", LastName: "lname", Identitiy: "identitiy", Password: "password"})
	// mock.ExpectBegin()
	// mock.ExpectExec("INSERT INTO users").WithArgs("name", "lname", "identitiy", "password", date.GetNowString()).WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()

	// // now we execute our method
	// err = Model.Register()
	// if err != nil {
	// 	t.Errorf("error was not expected while updating stats: %s", err)
	// }

	// // we make sure that all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}
