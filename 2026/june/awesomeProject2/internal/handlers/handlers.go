package handlers

import (
	"log"
	"main/internal/db"
	"main/internal/models"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	DB *db.Database
}

func NewEndpoint(Db *db.Database) *Endpoint {
	return &Endpoint{
		DB: Db,
	}
}
func (e *Endpoint) GetAllTracks(ctx *gin.Context) {
	var tracks []models.Track
	e.DB.DB.Find(&tracks)
	ctx.JSON(200, gin.H{"data": tracks})
}
func (e *Endpoint) CreateTrack(ctx *gin.Context) {
	var input models.Track
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		log.Println(err.Error())
		return
	}
	track := models.Track{Artist: input.Artist, Title: input.Title}
	e.DB.DB.Create(&track)
	ctx.JSON(200, gin.H{"data": input})
}
func (e *Endpoint) GetTrackById(ctx *gin.Context) {
	id := ctx.Param("id")
	var track models.Track
	e.DB.DB.First(&track, id)
	ctx.JSON(200, gin.H{"data": track})
}

func (e *Endpoint) DeleteTrackById(ctx *gin.Context) {
	id := ctx.Param("id")
	e.DB.DB.Delete(&models.Track{}, id)
	ctx.JSON(200, gin.H{"data": true})
}
