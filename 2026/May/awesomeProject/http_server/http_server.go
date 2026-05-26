package http_server

import (
	"fmt"
	"net/http"
)

func StartHttpServer() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from docker")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	return err
}
