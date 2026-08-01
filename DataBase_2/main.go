package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	db, err := NewDB("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to database")

	if err := Migrate(db); err != nil {
		log.Fatal(err)
	}

	// Dependency injection: main wires the concrete sqlite implementation
	// into the UserRepository interface that the rest of the program depends on.
	var repo UserRepository = NewSQLiteUserRepository(db)

	ctx := context.Background()

	newUsers := []struct {
		name, email, avatar string
	}{
		{"Kweks", "kweks@example.com", "http://avatar.com/kweks"},
		{"James", "james@example.com", "http://avatar.com/james"},
		{"John", "john@example.com", "http://avatar.com/john"},
		{"Joseph", "joseph@example.com", "http://avatar.com/joseph"},
	}

	for _, u := range newUsers {
		userID, err := repo.CreateUser(ctx, u.name, u.email, "password", u.avatar, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created user:", userID)
	}

	kweks, err := repo.GetUserByEmail(ctx, "kweks@example.com")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := json.MarshalIndent(kweks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Kweks:", string(bs))

	users, err := repo.GetUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total users:", len(users))
}
