package handlers

import (
	"Rest-API-NewsApp/config"
	"Rest-API-NewsApp/models"

	"github.com/gofiber/fiber/v2"
)

//Instalasi GORM dan pelajari hal ini

func GetNews(c *fiber.Ctx) error {
	id := c.Params("id")
	db, _ := config.ConnectDB()

	if id != "" {
		var news models.News
		err := db.QueryRow("SELECT id, title, content, created_at FROM news WHERE id = ?", id).Scan(
			&news.ID,
			&news.Title,
			&news.Content,
			&news.CreatedAt,
		)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "tidak dapat mengambil data id",
			})
		}

		return c.JSON(news)
	}

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
	result, err := db.Exec("INSERT INTO news (title, content) VALUES (?, ?)", news.Title, news.Content)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed insert data",
		})
	}

	id, _ := result.LastInsertId() //mengambil data id terakhir

	var CreatedNews models.News
	err = db.QueryRow("SELECT id, title, content, created_at FROM news WHERE id = ?", id).Scan(
		&CreatedNews.ID,
		&CreatedNews.Title,
		&CreatedNews.Content,
		&CreatedNews.CreatedAt,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch created data",
		})
	}

	return c.Status(201).JSON(CreatedNews)
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

	var updatedNews models.News
	err := db.QueryRow("SELECT id, title, content, created_at FROM news WHERE id = ?", id).Scan(
		&updatedNews.ID,
		&updatedNews.Title,
		&updatedNews.Content,
		&updatedNews.CreatedAt,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch update news",
		})
	}

	return c.Status(200).JSON(updatedNews)
}
