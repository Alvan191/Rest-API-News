package handlers

import (
	"Rest-API-NewsApp/config"
	"Rest-API-NewsApp/models"

	"github.com/gofiber/fiber/v2"
)

//Instalasi GORM dan pelajari hal ini

func GetNews(c *fiber.Ctx) error {
	db, _ := config.ConnectDB()
	rows, _ := db.Query("SELECT id, title, content, created_at FROM news")

	var news []models.News
	for rows.Next() {
		var n models.News
		rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt)
		news = append(news, n)
	}

	return c.JSON(news)
}

func CreateNews(c *fiber.Ctx) error {
	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	db, _ := config.ConnectDB()
	db.Exec("INSERT INTO news (title, content) VALUES (?, ?)", news.Title, news.Content)

	return c.Status(201).JSON(fiber.Map{
		"message": "News created successfully",
	})
}

func UpdateNews(c *fiber.Ctx) error {
	id := c.Params("id") // ✅ ambil id dari URL

	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	db, _ := config.ConnectDB()
	result, _ := db.Exec("UPDATE news SET title = ?, content = ? WHERE id = ?", news.Title, news.Content, id)

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "News not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "News update successfully",
	})
}
