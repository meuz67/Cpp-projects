package main

import (
	"awesomeProject3/featurePostgres/simple_connection"
	"awesomeProject3/featurePostgres/simple_sql"
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	task := simple_sql.TaskModel{Title: "Other title",
		Description: "Other description",
		CreatedAt:   time.Now(),
	}
	simple_sql.InsertRow(ctx, conn, task)
}
