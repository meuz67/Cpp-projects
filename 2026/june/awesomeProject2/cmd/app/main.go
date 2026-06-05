package main

import (
	"log"
	"main/internal/pkg"
)

func main() {
	a, err := pkg.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	a.Run(":8080")
}
