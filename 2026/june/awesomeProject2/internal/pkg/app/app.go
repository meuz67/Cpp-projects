package app

import (
	"awesomeProject2/internal/MW"
	"awesomeProject2/internal/app/endpoint"
	"awesomeProject2/internal/app/service"
	"fmt"

	"github.com/labstack/echo/v4"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func NewApp() *App {
	return &App{
		endpoint: endpoint.NewEndpoint(),
		service:  service.NewService(),
		echo:     echo.New(),
	}
}
func (app *App) Start() {
	app.echo.GET("/status", app.endpoint.Status)
	app.echo.Use(MW.CheckRole)
	if err := app.echo.Start(":8080"); err != nil {
		panic(err)
	}
	fmt.Println("Service started")
}
