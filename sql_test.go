package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// // a successful case
// func TestShouldUpdateStats(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	// now we execute our method
// 	if err = recordStats(db, 2, 3); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// // a failing test case
// func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec("INSERT INTO product_viewers").
// 		WithArgs(2, 3).
// 		WillReturnError(fmt.Errorf("some error"))
// 	mock.ExpectRollback()

// 	// now we execute our method
// 	if err = recordStats(db, 2, 3); err == nil {
// 		t.Errorf("was expecting an error, but there was none")
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

func TestCreateUser(t *testing.T) {
	user := User{Name: "AnuchitO", Age: 19}
	db, mock, _ := sqlmock.New()

	rows := sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(0, "AnuchitO", 19)

	// mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO users").WithArgs("AnuchitO", 19).WillReturnRows(rows)

	result, err := CreateUsers(db, user)
	assert.Nil(t, err)
	assert.EqualValues(t, result.Name, user.Name)
	assert.EqualValues(t, result.Age, user.Age)

}

func TestFindAllUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(0, "SomeUser0", 0).AddRow(0, "SomeUser1", 1).AddRow(0, "SomeUser2", 2)

	// mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, age FROM USERS").WillReturnRows(rows)

	result, err := FindAllUsers(db)
	assert.Nil(t, err)
	assert.EqualValues(t, len(result), 3)
}
