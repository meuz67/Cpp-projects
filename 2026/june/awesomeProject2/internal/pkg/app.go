package pkg

import (
	"main/internal/db"
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

type App struct {
	Database *db.Database
	Endpoint *handlers.Endpoint
	Router   *gin.Engine
}

func NewApp() (*App, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}

	endpoint := handlers.NewEndpoint(database)
	router := gin.Default()
	return &App{
		Database: database,
		Endpoint: endpoint,
		Router:   router,
	}, nil
}

func (a *App) Run(addr string) error {
	a.Router.GET("/tracks", a.Endpoint.GetAllTracks)
	a.Router.POST("/tracks", a.Endpoint.CreateTrack)
	a.Router.GET("/tracks/:id", a.Endpoint.GetTrackById)
	a.Router.DELETE("/tracks/:id", a.Endpoint.DeleteTrackById)
	return a.Router.Run(addr)
}
