package database

import (
	"github.com/dev-hyunsang/clone-todo-mate/ent"

	_ "github.com/mattn/go-sqlite3"
)

// DataBase Connection + AutoMigrations
func ConnectionSQLite() (*ent.Client, error) {
	client, err := ent.Open(
		"sqlite3",
		"file:clone-todo-mate.db?_fk=1")

	return client, err
}
