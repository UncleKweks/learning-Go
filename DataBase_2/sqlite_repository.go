package main

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// SQLiteUserRepository is the sqlite-backed implementation of UserRepository.
type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

var _ UserRepository = (*SQLiteUserRepository)(nil)

// CreateUser inserts a user and their profile in a single transaction so the
// two rows either both land or neither does.
func (r *SQLiteUserRepository) CreateUser(ctx context.Context, name, email, hashedPassword, avatar string, createdAt time.Time) (int64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password, created_at) VALUES (?, ?, ?, ?);`)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	res, err := stmt.Exec(name, email, string(hp), createdAt)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	profileStmt, err := tx.PrepareContext(ctx, `INSERT INTO profiles (user_id, avatar) VALUES(?,?)`)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer profileStmt.Close()

	if _, err := profileStmt.Exec(userID, avatar); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *SQLiteUserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	statement := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?;`

	row := r.db.QueryRowContext(ctx, statement, email)
	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPw, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SQLiteUserRepository) GetUsers(ctx context.Context) ([]User, error) {
	statement := `SELECT id, name, email, hashed_password, created_at FROM users`

	rows, err := r.db.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPw, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
