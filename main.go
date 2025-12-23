package main

import (
	"Rest-API-NewsApp/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/news", handlers.GetNews)
	app.Post("/news", handlers.CreateNews)
	app.Put("/news/:id", handlers.UpdateNews)

	log.Fatal(app.Listen(":8080"))
}
