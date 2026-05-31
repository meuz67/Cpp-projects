package repository

import (
	"encoding/json"
	"io"
	"slices"

	"awesomeProject6/internal/models"
)

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{},
	}
}

func (r *UserRepository) Create(user models.User) models.User {
	r.users = append(r.users, user)
	return user
}

func (r *UserRepository) GetAll() []models.User {
	return r.users
}

func (r *UserRepository) UpdateByID(id string, body io.Reader) (models.User, bool) {
	for i, user := range r.users {
		if user.ID == id {
			json.NewDecoder(body).Decode(&r.users[i])
			return r.users[i], true
		}
	}
	return models.User{}, false
}

func (r *UserRepository) DeleteByID(id string) bool {
	for i, user := range r.users {
		if user.ID == id {
			r.users = slices.Delete(r.users, i, i)
			return true
		}
	}
	return false
}
