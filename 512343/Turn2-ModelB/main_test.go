package main

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// Define the callback function signature
type TransactionalCallback func(tx *sql.Tx) error

// Function to perform a transaction and execute the callback
func PerformTransaction(db *sql.DB, callback TransactionalCallback) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Re-panic to allow testing to handle it
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = callback(tx)
	return err
}

// Example callback function that inserts data into the "users" table
func InsertUser(tx *sql.Tx, userName string) error {
	stmt, err := tx.Prepare("INSERT INTO users (username) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName)
	return err
}

// Function to check if a user exists in the database
func UserExists(db *sql.DB, userName string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", userName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Test setup and teardown for SQLite database
func setupDatabase(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Error opening SQLite database: %v", err)
	}

	// Create the "users" table
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, username TEXT)")
	if err != nil {
		t.Fatalf("Error creating table: %v", err)
	}

	return db
}

func teardownDatabase(t *testing.T, db *sql.DB) {
	err := db.Close()
	if err != nil {
		t.Errorf("Error closing SQLite database: %v", err)
	}
}


func TestTransactions(t *testing.T) {
	tests := []struct {
		name      string
		callback  TransactionalCallback
		userName  string
		wantExist bool
		wantErr   bool
	}{
		{
			name: "Successful Transaction",
			callback: func(tx *sql.Tx) error {
				return InsertUser(tx, "testuser")
			},
			userName:  "testuser",
			wantExist: true,
			wantErr:   false,
		},
		{
			name: "Rollback on Failure",
			callback: func(tx *sql.Tx) error {
				err := InsertUser(tx, "testuser")
				if err != nil {
					return err
				}
				return errors.New("some other error")
			},
			userName:  "testuser",
			wantExist: false,
			wantErr:   true,
		},