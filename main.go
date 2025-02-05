package main

import (
	"os"

	"github.com/AKAZJAYA/blogbackend/database"
	"github.com/AKAZJAYA/blogbackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {

	database.Connect()
	
	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	app := fiber.New()

	routes.Setup(app)
	app.Listen(":" + port)
}