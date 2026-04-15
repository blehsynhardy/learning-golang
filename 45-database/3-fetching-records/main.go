package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

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

	// 	lastId, err := insertData(db, "Awodele Esther", "esther@gmail.com", "password2$")

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("last user id is", lastId);

	// 	lastId, err = insertData(db, "Awodele boluwatife", "boluwatife@gmail.com", "password2$")

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("last user id is", lastId);

	// 	lastId, err = insertData(db, "Awodele deborah", "deborahagesin@gmail.com", "password2$")

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("last user id is", lastId);

	// 	lastId, err = insertData(db, "Awodele adedayo", "adedayo@gmail.com", "password2$")

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("last user id is", lastId)

	janet, err := getUserByEmail(db, "esther@gmail.com")

	if err != nil {
		log.Fatal(err)
	}

	if janet == nil {
		fmt.Println("user not found")
	} else {
		fmt.Printf("user found: %v\n", janet)
	}

	janetJson, err := json.Marshal(janet)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user in json format: %s\n", janetJson)


	users, err := getAllUsers(db)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all users: %v\n", users)
}

func createTable(db *sql.DB) error {

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func getUserByEmail(db *sql.DB, email string) (*User, error) {

	stmt := `SELECT id, name, email, password, createdAt FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)

	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
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

func getAllUsers(db *sql.DB) ([]User, error) {

	stmt := `SELECT id, name, email, password, createdAt FROM users`

	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
