package main

import (
	"log"

	"github.com/Jiran03/mailku/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := routes.New()

	// e.Start(":8080")
	e.Logger.Fatal(e.Start(":8080"))
}
