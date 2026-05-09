package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}
type Users struct {
	Users []User
}

func handler(w http.ResponseWriter, r *http.Request) {
	u1 := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}
	u2 := User{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       45,
	}
	u3 := User{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       67,
	}
	data := Users{
		Users: []User{u1, u2, u3},
	}
	tmp := template.Must(template.ParseFiles("index.html"))
	tmp.Execute(w, data)
}
func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
