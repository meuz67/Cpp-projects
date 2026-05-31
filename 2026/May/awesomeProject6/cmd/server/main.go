package main

import (
	"log"
	"net/http"

	"awesomeProject6/internal/handler"
	"awesomeProject6/internal/repository"
	"awesomeProject6/internal/router"
)

func main() {
	userRepo := repository.NewUserRepository()

	homeHandler := handler.NewHomeHandler()
	userHandler := handler.NewUserHandler(userRepo)

	r := router.New(homeHandler, userHandler)

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
