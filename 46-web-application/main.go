package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/46-web-application/repository"
	"os"

	"github.com/golangcollege/sessions"

	_ "modernc.org/sqlite"
)

type Application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    repository.UserRepository
	templateDir string
	rc          *RenderCache
	publicDir   string
	session     *sessions.Session
}

func main() {

	// In a production application, you would want to use a more secure key and store it in an environment variable or a configuration file.
	//this is just for demonstration purposes and how to use sessions in a web application. In a real application, you should use a secure key and store it securely.

	var secret = []byte("3k4l5m6n7o8p9q0r1s2t3u4v5w6x7y8z9a0b1c2d3e4f5g6h7i8j9k0l1m2n3o4p5q6r7s8t")

	session := sessions.New(secret)
	session.Lifetime = 12 * 60 * 60 //3 hours
	session.Secure = false // dev server runs plain HTTP; Secure cookies are dropped by browsers without TLS
	session.HttpOnly = true
	// session.SameSite = http.SameSiteStrictMode

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
		session:     session,
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

	// _, err = db.Exec(`ALTER TABLE profile ADD COLUMN userId INT NOT NULL`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS users (
	// 		id    INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		name  TEXT    NOT NULL,
	// 		email TEXT    NOT NULL UNIQUE,
	// 		hashed_password TEXT NOT NULL,
	// 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	// 		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	// 	)
	// `)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//  p.bio, p.avatar, p.createdAt, p.updatedAt

	// _, err = db.Exec(
	// 	`CREATE TABLE profile (id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	bio VARCHAR(1000) NOT NULL,
	// 	avatar TEXT NOT NULL,
	// 	created_at TIMESTAMP NOT NULL DEFAULT CURRRENT_TIMESTAMP,
	// 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

	// 	)`)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// _, err = db.Exec(`
	// 	insert into users (name, email, hashed_password) values
	// 	('Joseph', 'joseph@example.com', '$2a$12$WLT7jewOJk31nEa4SxjslOcfXkUCO.nshqhSbRF23pB7Og51SbSM6')
	// `)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Table created")
	fmt.Println("Databse Connection successful")

	// _, err = db.Exec(schema)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("table created successfully")
	return db, nil
}
