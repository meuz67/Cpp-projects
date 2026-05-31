package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"awesomeProject6/internal/models"
	"awesomeProject6/internal/repository"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user = h.repo.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.repo.GetAll())
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user, ok := h.repo.UpdateByID(id, r.Body)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if h.repo.DeleteByID(id) {
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
