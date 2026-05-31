package handler

import (
	"encoding/json"
	"net/http"

	"awesomeProject6/internal/models"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Response{
		Message: "Welcome to the home page",
	})
}
