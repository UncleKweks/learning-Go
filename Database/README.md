# Database

A small exercise connecting to a SQLite database (via `modernc.org/sqlite`), demonstrating table creation, inserts with `bcrypt`-hashed passwords, and querying a single row vs. all rows.

## What it does

- Creates a `users` table if it doesn't already exist.
- Inserts four users, hashing each password with `bcrypt` before storing it.
- Fetches a single user by email (`GetUserByEmail`) and prints it as JSON.
- Inserts an additional user using a prepared statement (`createUserWithPrepared`).
- Fetches all users (`GetUsers`) and prints the full list as JSON (currently commented out).

## Running

```bash
go run ./Database
```

The database file (`users_database.db`) is recreated on each run and is git-ignored.
