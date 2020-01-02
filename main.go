package main

import (
	"go_framework/routes"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := routes.Index()

	// Start server
	e.Logger.Fatal(e.Start(":8000"))

}
