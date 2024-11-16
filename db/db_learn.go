package db_learn

import (
	"context"

	"encore.dev/storage/sqldb"
)

var db_learn_db = sqldb.NewDatabase("db_learn", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

func insert(ctx context.Context, id, title string, done bool) error {
	_, err := db_learn_db.Exec(ctx, `
		INSERT INTO db_learn_tb (id, title, done)
		VALUES ($1, $2, $3)
	`, id, title, done)
	return err
}
