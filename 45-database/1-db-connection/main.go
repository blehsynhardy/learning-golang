package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "modernc.org/sqlite"
)





func main() {


	dbName := "data.db"

	_ = os.Remove(dbName)

db, err := sql.Open("sqlite", dbName)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {

		fmt.Println("Closing db connection")
		
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()


	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Databse Connection successful");

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id    INTEGER PRIMARY KEY AUTOINCREMENT,
			name  TEXT    NOT NULL,
			email TEXT    NOT NULL UNIQUE
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created")






}


