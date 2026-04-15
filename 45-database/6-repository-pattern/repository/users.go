package repository

import (
	"context"
	"database/sql"
	"main/45-database/6-repository-pattern/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUserWithProfile(name, email, password, bio, avatar string) (int64, error)
	FetchAllUsers() ([]models.User, error)
	FetchUserByEmail(email string) (*models.User, error)
}

type SqliteUserRepository struct {
	db *sql.DB
}

func NewSqliteUserRepository(db *sql.DB) UserRepository {
	return &SqliteUserRepository{db: db}
}

func (r *SqliteUserRepository) CreateUserWithProfile(name, email, password, bio, avatar string) (int64, error) {

	hsp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
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

func (r *SqliteUserRepository) FetchAllUsers() ([]models.User, error) {

	ctx := context.Background()
	query := `
        SELECT u.id, u.name, u.email, u.password, u.createdAt,
               p.bio, p.avatar, p.createdAt, p.updatedAt
        FROM users u
        JOIN profiles p ON u.id = p.userId`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
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

func (r *SqliteUserRepository) FetchUserByEmail(email string) (*models.User, error) {

	ctx := context.Background()

	query := `SELECT u.id, u.name, u.email, u.password, u.createdAt,
			   p.bio, p.avatar, p.createdAt, p.updatedAt
		FROM users u
		JOIN profiles p ON u.id = p.userId
		WHERE u.email = ?`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, query, email)

	var user models.User

	err = row.Scan(
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

	return &user, nil

}
