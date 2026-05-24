package main

import (
	"awesomeProject3/featurePostgres/simple_connection"
	"awesomeProject3/featurePostgres/simple_sql"
	"context"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	tasks, _ := simple_sql.SelectRows(ctx, conn)
	if tasks != nil {
		for _, v := range tasks {
			pp.Println(v)
		}
	}
}
