package simple_action

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func UpdateTask(ctx context.Context, db *pgx.Conn, id int) error {
	_, err := db.Exec(ctx, `UPDATE todo SET is_done=true, done_at=$1 WHERE id = $2`, time.Now(), id)
	return err
}
