package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sqlx.DB
}

func New(dbName string) *DB {
	rootDir, _ := os.Getwd()
	db, err := sqlx.Connect("sqlite3", filepath.Join(rootDir, dbName))
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db}
}

func (db *DB) RunMigrations() error {
	sqlCreate := `
	CREATE TABLE IF NOT EXISTS accounts (id INTEGER PRIMARY KEY, username TEXT, password TEXT);
	`

	_, err := db.Exec(sqlCreate)
	return err
}

func Teardown(dbName string) {
	rootDir, _ := os.Getwd()
	err := os.Remove(filepath.Join(rootDir, dbName))
	if err != nil {
		log.Fatal(err)
	}
}
