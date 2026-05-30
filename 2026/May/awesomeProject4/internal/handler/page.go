package handler

import (
	"app/internal/models"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	Title        string
	Flash        string
	Error        string
	Username     string
	LoggedIn     bool
	Task         *models.Task
	Tasks        []models.Task
	FormUsername string
}

func pageFromContext(c echo.Context, title string) PageData {
	p := PageData{Title: title}
	if name, ok := c.Get("username").(string); ok && name != "" {
		p.Username = name
		p.LoggedIn = true
	}
	return p
}
