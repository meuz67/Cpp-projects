package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, "postgres://postgres:sar58yeaf@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to postgres")
}
