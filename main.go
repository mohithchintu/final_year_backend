package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/final_year_backend/handlers"
)

func main() {
	app := fiber.New()

	app.Post("/generate", handlers.GenerateDevices)
	app.Post("/authenticate", handlers.AuthenticateDevices)

	log.Fatal(app.Listen(":5005"))
}
