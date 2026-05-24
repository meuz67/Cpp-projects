package simple_action

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteTask(ctx context.Context, db *pgx.Conn, id int) error {
	_, err := db.Exec(ctx, `DELETE FROM todo WHERE id = $1`, id)
	return err
}
