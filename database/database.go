package database

import (
	"context"
	"github.com/dev-hyunsang/clone-todo-mate/ent"
)

// DataBase Connection + AutoMigrations
func ConnectionSQLite() (*ent.Client, error) {
	client, err := ent.Open(
		"sqlite3",
		"file:clone-todo-mate.db?_fk=1")

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, err
}
