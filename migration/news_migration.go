package migration

import (
	"Rest-API-NewsApp/config"
	"Rest-API-NewsApp/models"
)

func MigrateNewsAuthor() {
	db := config.DB

	if !db.Migrator().HasColumn(&models.News{}, "Author") {
		db.Migrator().AddColumn(&models.News{}, "Author")
	}
}
