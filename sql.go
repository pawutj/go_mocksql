package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func CreateUsers(db *sql.DB, user User) (User, error) {

	row := db.QueryRow(`INSERT INTO users (name, age) values ($1, $2)  RETURNING id,name,age`, user.Name, user.Age)

	result := User{}
	err := row.Scan(&result.ID, &result.Name, &result.Age)
	if err != nil {
		return User{}, err
	}

	return result, nil
}

func FindAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age FROM USERS")

	result := []User{}

	if err != nil {
		return result, err
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return result, err
		}
		result = append(result, user)
	}
	return result, nil
}

func FindOneUser(db *sql.DB, id int) (User, error) {
	result := User{}

	stmt, err := db.Prepare("SELECT id, name, age FROM users where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
		return result, err
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&result.ID, &result.Name, &result.Age)

	if err != nil {
		log.Fatal("can't Scan row into variables", err)
		return result, err
	}

	return result, nil

}
func CreateTable(db *sql.DB) error {
	createTb := `
	CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT );
	`

	_, err := db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	return nil
}
