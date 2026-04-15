package main

// In this example, we will learn how to use transactions in Go with a SQLite database. Transactions allow us to execute a series of database operations as a single unit of work. If any operation within the transaction fails, the entire transaction can be rolled back, ensuring data integrity.

//to do
//1.User create account
//2. Create Wallet for user
//3. Add money to wallet
//4. Write a transaction log for the wallet transaction

import (
	"context"
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
var profileSchema = `
	CREATE TABLE IF NOT EXISTS profiles (
		userId INTEGER PRIMARY KEY NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		bio TEXT,
		avatar TEXT,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
`

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Profile   Profile   `json:"profile"`
}

type Profile struct {
	UserID    int       `json:"user_id"`
	Bio       string    `json:"bio"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {

	dbName := "tiklog.db"
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

	err = createTables(db)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tables created successfully")

	// userID, err := createUserWithProfile(db, "Awodele Esther", "awodelee@example.com", "password123", "Software Engineer", "avatar.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("User created with ID: %d\n", userID)

	users, err := fetchAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("All users:")
	// for _, user := range users {
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s, Password: %s, Created At: %s\n",
	// 		user.ID, user.Name, user.Email, user.Password, user.CreatedAt)
	// 	fmt.Printf("Profile - Bio: %s, Avatar: %s, Created At: %s, Updated At: %s\n",
	// 		user.Profile.Bio, user.Profile.Avatar, user.Profile.CreatedAt, user.Profile.UpdatedAt)
	// }

	jsonUsers, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All users in JSON format:")
	fmt.Println(string(jsonUsers))

}

func createTables(db *sql.DB) error {
	_, err := db.Exec(schema)

	if err != nil {
		return err
	}

	_, err = db.Exec(profileSchema)

	if err != nil {
		return err
	}

	return nil
}

func createUserWithProfile(db *sql.DB, name, email, password, bio, avatar string) (int64, error) {

	hsp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	userStmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, password) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer userStmt.Close()

	res, err := userStmt.ExecContext(ctx, name, email, string(hsp))
	if err != nil {
		return 0, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	profileStmt, err := tx.PrepareContext(ctx, `INSERT INTO profiles (userId, bio, avatar) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer profileStmt.Close()

	_, err = profileStmt.ExecContext(ctx, userID, bio, avatar)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func fetchAllUsers(db *sql.DB) ([]User, error) {

	ctx := context.Background()
	query := `
        SELECT u.id, u.name, u.email, u.password, u.createdAt,
               p.bio, p.avatar, p.createdAt, p.updatedAt
        FROM users u
        JOIN profiles p ON u.id = p.userId`

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.Profile.Bio,
			&user.Profile.Avatar,
			&user.Profile.CreatedAt,
			&user.Profile.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
