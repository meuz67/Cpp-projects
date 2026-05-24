package simple_action

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertInto(ctx context.Context, db *pgx.Conn, task Task) error {
	_, err := db.Exec(ctx, "INSERT INTO todo(name, description, is_done, created_at, done_at) VALUES($1, $2, $3, $4, $5)", task.Name, task.Description, false, time.Now(), nil)
	return err
}
