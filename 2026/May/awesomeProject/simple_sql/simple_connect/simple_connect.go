package simple_connect

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, os.Getenv("connStr"))
}
