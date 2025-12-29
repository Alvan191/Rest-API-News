package main

import (
	"Rest-API-NewsApp/config"
	"Rest-API-NewsApp/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	app.Get("/news", handlers.GetNews)
	app.Get("/news/:id", handlers.GetNews)
	app.Post("/news", handlers.CreateNews)
	app.Put("/news/:id", handlers.UpdateNews)
	app.Delete("/news/:id", handlers.DeleteNews)

	log.Fatal(app.Listen(":8080"))
}
