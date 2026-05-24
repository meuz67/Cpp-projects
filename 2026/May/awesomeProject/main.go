package main

import (
	"awesomeProject/simple_sql/simple_action"
	"awesomeProject/simple_sql/simple_connect"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	db, err := simple_connect.Connect(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}
	defer db.Close(ctx)
	err = simple_action.DeleteTask(ctx, db, 2)
	if err != nil {
		panic(err)
	}
}
