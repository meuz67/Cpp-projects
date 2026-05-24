package simple_action

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, db *pgx.Conn) {
	db.Exec(ctx, `CREATE TABLE IF NOT EXISTS todo(
    id serial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    is_done BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
	done_at TIMESTAMP 
    )`)
}
