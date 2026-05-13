package httpServer

import (
	"errors"
	"net/http"
)

func StartHttpServer() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Docker"))
	})
	err := http.ListenAndServe(":8081", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
