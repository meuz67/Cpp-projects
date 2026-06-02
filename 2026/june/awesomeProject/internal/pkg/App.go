package pkg

import (
	"june/internal/MW"
	"june/internal/endpoint"
	"june/internal/service"
	"log"

	"github.com/labstack/echo/v4"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func NewApp() *App {
	app := &App{}
	app.service = service.NewService()
	app.endpoint = endpoint.New(app.service)
	app.echo = echo.New()
	app.echo.GET("/status", app.endpoint.Status)
	app.echo.Use(MW.CheckRole)
	return app
}
func (app *App) Start() error {
	log.Println("Server starting...")
	if err := app.echo.Start(":8080"); err != nil {
		return err
	}
	return nil
}
