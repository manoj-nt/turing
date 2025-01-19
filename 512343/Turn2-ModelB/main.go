package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ExecuteInTransaction(db *sql.DB, callback func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := callback(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func InsertUser(tx *sql.Tx, username string) error {
	_, err := tx.Exec("INSERT INTO users (username) VALUES (?);", username)
	return err
}
