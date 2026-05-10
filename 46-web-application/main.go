package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/46-web-application/repository"
	"os"

	_ "modernc.org/sqlite"
)

type Application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    repository.UserRepository
	templateDir string
	rc          *RenderCache
	publicDir   string
}

func main() {

	db, err := connectToDatabase("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &Application{
		errorLog:    log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile),
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile),
		userRepo:    repository.NewSqliteUserRepository(db),
		templateDir: "./templates",
		publicDir:   "./public",
	}

	app.rc = NewRenderCache(true, app.templateDir)

	log.Print("Application initialized successfully running on port 4000")

	if err := app.serve(); err != nil {
		log.Fatal()
	}
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
