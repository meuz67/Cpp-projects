package db

import (
	"log"
	"main/internal/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {
	conn, err := Connect()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Database{
		DB: conn,
	}, nil
}
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Track{})
	return db, nil
}
