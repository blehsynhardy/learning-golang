package main

// In this example, we will learn how to implement the repository pattern in Go with a SQLite database. The repository pattern is a design pattern that abstracts the data access layer, allowing us to separate the business logic from the data access logic. This makes our code more modular, testable, and maintainable.

// We will create a simple user management system where we can create users with profiles and fetch user data from the database. We will define a UserRepository interface that will have methods for creating a user with a profile, fetching all users, and fetching a user by email. We will then implement this interface using a SQLite database.

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/45-database/6-repository-pattern/repository"

	_ "modernc.org/sqlite"
)

var schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);



CREATE TABLE IF NOT EXISTS profiles (
	userId INTEGER PRIMARY KEY,
	bio TEXT,
	avatar TEXT,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);
`

func main() {
	dbName := "tikloger.db"
	db, err := connectToDatabase(dbName)
	checkError(err)
	defer db.Close()

	fmt.Println("connection successsful")

	repo := repository.NewSqliteUserRepository(db)

	//create user with profile
	userID, err := repo.CreateUserWithProfile("Awodele Repository3", "awodele122e@example.com", "password123", "Software Engineer", "avatar.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User created with ID: %d\n", userID)

	//fetch all users
	printUser(repo)

}

func printUser(repo repository.UserRepository) {
	users, err := repo.FetchAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	jsonUsers, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("All users dependency: %s\n", jsonUsers)
}

func connectToDatabase(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Databse Connection successful")

	// _, err = db.Exec(schema)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("table created successfully")
	return db, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
