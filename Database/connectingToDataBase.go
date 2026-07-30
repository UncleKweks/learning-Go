package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "modernc.org/sqlite"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	HashedPw  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	dbName := "users_database.db"
	if err := os.Remove(dbName); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable(db)

	lastID, err := createUserTable(db, "Kweks", "kweks@example.com", "password", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last ID:", lastID)

	lastID, err = createUserTable(db, "James", "James@example.com", "password", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last ID:", lastID)

	lastID, err = createUserTable(db, "John", "John@example.com", "password", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last ID:", lastID)

	lastID, err = createUserTable(db, "Joseph", "Joseph@example.com", "password", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last ID:", lastID)

	kweks, err := GetUserByEmail(db, "kweks@example.com")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := json.MarshalIndent(kweks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Kweks:", string(bs))

	users, err := GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	usersJSON, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Users:", string(usersJSON))
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserTable(db *sql.DB, name, email, hashedPassword string, createdAt time.Time) (int64, error) {
	statement := `INSERT INTO users (name, email, hashed_password, created_at) VALUES (?, ?, ?, ?);`

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(statement, name, email, string(hp), createdAt)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	statement := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?;`

	row := db.QueryRow(statement, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPw, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	statement := `SELECT id, name, email, hashed_password, created_at FROM users`
	rows, err := db.Query(statement)
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
