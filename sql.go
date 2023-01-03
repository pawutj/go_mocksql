package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
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

func main() {
	// @NOTE: the real connection is not required for tests
	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = recordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
		panic(err)
	}
}
