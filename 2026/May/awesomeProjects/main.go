package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Product struct {
	Id      int
	Model   string
	Company string
	Price   int
}

var db *sql.DB

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Exec("DELETE FROM products WHERE id = $1", id)
	http.Redirect(w, r, "/", http.StatusFound)
}
func editPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product := Product{}
	err = db.QueryRow("SELECT * FROM products WHERE id=$1", id).Scan(&product.Id,
		&product.Model, &product.Company, &product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmp, err := template.ParseFiles("edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmp.Execute(w, product)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	model := r.FormValue("model")
	company := r.FormValue("company")
	price := r.FormValue("price")
	db.Exec("UPDATE products SET model=$1, company=$2, price=$3 WHERE id=$4", model, company, price, id)
	http.Redirect(w, r, "/", http.StatusFound)
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmp, err := template.ParseFiles("create.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmp.Execute(w, nil)
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		model := r.PostFormValue("model")
		company := r.PostFormValue("company")
		price := r.PostFormValue("price")
		db.Exec("INSERT INTO products (model, company, price) VALUES ($1, $2, $3)", model, company, price)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")
	query := "SELECT * FROM products"
	switch sort {
	case "price_asc":
		query += " ORDER BY price ASC"
		break
	case "price_desc":
		query += " ORDER BY price DESC"
		break
	}
	result, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()
	var products []Product
	for result.Next() {
		p := Product{}
		err := result.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err.Error())
		}
		products = append(products, p)
	}
	tmp1 := template.Must(template.ParseFiles("index.html"))
	tmp1.Execute(w, products)
}
func main() {
	connStr := "user=postgres password=sar58yeaf dbname=postgres sslmode=disable"
	var err error
	rout := mux.NewRouter()
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	rout.HandleFunc("/", mainHandler)
	rout.HandleFunc("/create", createHandler)
	rout.HandleFunc("/edit/{id:[0-9]+}", editPage).Methods("GET")
	rout.HandleFunc("/edit/{id:[0-9]+}", editHandler).Methods("POST")
	rout.HandleFunc("/delete/{id:[0-9]+}", deleteHandler)
	http.Handle("/", rout)
	http.ListenAndServe(":8080", nil)
}
