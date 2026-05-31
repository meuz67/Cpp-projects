package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"awesomeProject6/internal/handler"
)

func New(home *handler.HomeHandler, users *handler.UserHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", home.ServeHTTP).Methods("GET")
	r.HandleFunc("/users", users.Create).Methods("POST")
	r.HandleFunc("/users", users.GetAll).Methods("GET")
	r.HandleFunc("/users", users.Update).Methods("PUT")
	r.HandleFunc("/users", users.Delete).Methods("DELETE")
	return r
}
