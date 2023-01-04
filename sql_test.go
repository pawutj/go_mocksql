package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := User{Name: "SomeUser", Age: 27}
	db, mock, _ := sqlmock.New()

	rows := sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(0, "SomeUser", 27)

	mock.ExpectQuery("INSERT INTO users").WithArgs("SomeUser", 27).WillReturnRows(rows)

	result, err := CreateUsers(db, user)
	assert.Nil(t, err)
	assert.EqualValues(t, result.Name, user.Name)
	assert.EqualValues(t, result.Age, user.Age)

}

func TestFindAllUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(0, "SomeUser0", 0).AddRow(0, "SomeUser1", 1).AddRow(0, "SomeUser2", 2)
	mock.ExpectQuery("SELECT id, name, age FROM USERS").WillReturnRows(rows)

	result, err := FindAllUsers(db)
	assert.Nil(t, err)
	assert.EqualValues(t, len(result), 3)
}

func TestFindOneUser(t *testing.T) {

	db, mock, _ := sqlmock.New()

	row := sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(0, "SomeUser", 0)
	mock.ExpectPrepare("SELECT id, name, age FROM users").ExpectQuery().WithArgs(0).WillReturnRows(row)
	result, err := FindOneUser(db, 0)

	assert.Nil(t, err)
	assert.EqualValues(t, result.Name, "SomeUser")

}

func TestCreateTable(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))

	err := CreateTable(db)
	assert.Nil(t, err)

}
