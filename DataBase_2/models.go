package main

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	HashedPw  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Profile   Profile   `json:"profile"`
}

type Profile struct {
	ID        int       `json:"id"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}
