package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Product struct {
	Id      int
	Model   string
	Company string
	Price   int
}

var db *sql.DB

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmp, err := template.ParseFiles("create.html")
		if err != nil {
			fmt.Println(err)
		}
		tmp.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		model := r.PostFormValue("model")
		company := r.PostFormValue("company")
		price, err := strconv.Atoi(r.PostFormValue("price"))
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		db.Exec(`INSERT INTO products (model, company, price) VALUES ($1, $2, $3)`, model, company, price)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	defer rows.Close()
	users := make([]Product, 0)
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		users = append(users, p)
	}
	tmp1 := template.Must(template.ParseFiles("index.html"))
	tmp1.Execute(w, users)
}
func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=sar58yeaf dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.ListenAndServe(":8080", nil)
}
