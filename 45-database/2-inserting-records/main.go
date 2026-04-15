package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "modernc.org/sqlite"
)

var schema = `
		CREATE TABLE IF NOT EXISTS users (
			id    INTEGER PRIMARY KEY AUTOINCREMENT,
			name  TEXT    NOT NULL,
			email TEXT    NOT NULL UNIQUE,
			password TEXT NOT NULL,
			createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

func main() {

	dbName := "users_database.db"

	db, err := sql.Open("sqlite", dbName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Databse Connection successful")

	//createTable(db)

	fmt.Println("table created successfully")

	lastId, err := insertData(db, "Awodele Esther", "esther@gmail.com", "password2$")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("last user id is", lastId);


	lastId, err = insertData(db, "Awodele boluwatife", "boluwatife@gmail.com", "password2$")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("last user id is", lastId);


	lastId, err = insertData(db, "Awodele deborah", "deborahagesin@gmail.com", "password2$")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("last user id is", lastId);


	lastId, err = insertData(db, "Awodele adedayo", "adedayo@gmail.com", "password2$")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("last user id is", lastId)
}

func createTable(db *sql.DB) error {

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func insertData(db *sql.DB, name, email, password string) (int64, error) {

	stmt := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	//hash pasasword

	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	res, err := db.Exec(stmt, name, email, string(hp))

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()

}
