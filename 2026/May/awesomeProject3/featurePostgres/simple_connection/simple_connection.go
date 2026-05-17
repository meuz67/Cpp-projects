package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(connectionString string) {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to postgres")
}
