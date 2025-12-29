package handlers

import (
	"Rest-API-NewsApp/config"
	"Rest-API-NewsApp/models"

	"github.com/gofiber/fiber/v2"
)

func GetNews(c *fiber.Ctx) error {
	id := c.Params("id")

	if id != "" {
		var news models.News
		result := config.DB.First(&news, id)

		if result.Error != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "tidak dapat mengambil data id",
			})
		}
		return c.JSON(news)
	}

	var news []models.News
	config.DB.Find(&news)
	return c.JSON(news)
}

func CreateNews(c *fiber.Ctx) error {
	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if news.Title == "" || news.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "data title atau content tidak boleh kosong",
		})
	}

	result := config.DB.Create(&news)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "gagal membuat data",
		})
	}

	return c.Status(201).JSON(news)
}

func UpdateNews(c *fiber.Ctx) error {
	id := c.Params("id") // ✅ ambil id dari URL

	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	if news.Title == "" && news.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "tidak ada data yang diupdate",
		})
	}

	config.DB.Model(&models.News{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":   news.Title,
			"content": news.Content,
		})

	var updatedNews models.News
	config.DB.First(&updatedNews, id)

	return c.Status(200).JSON(updatedNews)
}

func DeleteNews(c *fiber.Ctx) error {
	id := c.Params("id")

	result := config.DB.Delete(&models.News{}, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "news not found",
		})
	}

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete news",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "news deleted successfully",
	})
}
