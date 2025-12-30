package handlers

import (
	"Rest-API-NewsApp/config"
	"Rest-API-NewsApp/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

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

	if err := validate.Struct(news); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
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
	now := time.Now().UTC()
	id := c.Params("id")

	var news models.News

	if err := c.BodyParser(&news); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	if err := validate.Struct(news); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// bentuk menggunakan update bukan save
	updates := map[string]interface{}{
		"title":      news.Title,
		"content":    news.Content,
		"updated_at": &now,
	}
	config.DB.Model(&models.News{}).Where("id = ?", id).Updates(updates)

	config.DB.First(&news, id)
	return c.Status(200).JSON(news)
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
