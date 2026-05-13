package httpServer

import (
	"errors"
	"net/http"
)

func StartHttpServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from docker"))
	})
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
