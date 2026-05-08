package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Id   int
	Name string
}

func main() {
	connStr := "user=postgres password=sar58yeaf dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	products := []User{}
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
		products = append(products, u)
		fmt.Println(u.Id, " - ", u.Name)
	}
}
