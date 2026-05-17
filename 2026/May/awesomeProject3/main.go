package main

import (
	"awesomeProject3/featurePostgres/simple_connection"
	"os"
)

func main() {
	connStr := os.Getenv("connStr")
	simple_connection.CheckConnection(connStr)
}
