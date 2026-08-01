package main

import (
	"context"
	"time"
)

// UserRepository abstracts persistence for users so callers depend on this
// interface rather than a concrete database implementation.
type UserRepository interface {
	CreateUser(ctx context.Context, name, email, hashedPassword, avatar string, createdAt time.Time) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
}
