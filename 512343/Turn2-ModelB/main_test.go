package main

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, username TEXT);")
	if err != nil {
		t.Fatal(err)
	}

	teardown := func() {
		db.Close()
	}

	return db, teardown
}

func TestExecuteInTransaction(t *testing.T) {
	tests := []struct {
		name        string
		callback    func(*sql.Tx) error
		expectError bool
	}{
		{
			"Successful transaction",
			func(tx *sql.Tx) error {
				return InsertUser(tx, "john_doe")
			},
			false,
		},
		{
			"Rollback on error",
			func(tx *sql.Tx) error {
				err := InsertUser(tx, "john_doe")
				if err == nil {
					return errors.New("force rollback")
				}
				return nil
			},
			true,
		},
		{
			"Concurrency issue simulation",
			func(tx *sql.Tx) error {
				// Simulate a long-running transaction that could cause a concurrency issue
				_, err := tx.Exec("SELECT sleep(1);")
				return err
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := setupTestDB(t)
			defer teardown()

			err := ExecuteInTransaction(db, tt.callback)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error = %v, got %v", tt.expectError, err)
			}
		})
	}
}
