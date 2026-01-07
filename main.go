package main

import (
	"Rest-API-NewsApp/config"
	"Rest-API-NewsApp/handlers"
	"Rest-API-NewsApp/migration"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	migration.MigrateNewsAuthor()

	app := fiber.New()

	app.Get("/news", handlers.GetNews)
	app.Get("/news/:id", handlers.GetNews)
	app.Post("/news", handlers.CreateNews)
	app.Put("/news/:id", handlers.UpdateNews)
	app.Delete("/news/:id", handlers.DeleteNews)

	log.Fatal(app.Listen(":8080"))
}

// type Comments struct {
// 	ID        uint      `gorm:"primaryKey" json:"id"`
// 	UserID    uint      `gorm:"not null" json:"user_id"`
// 	PostID    uint      `gorm:"not null" json:"post_id"`
// 	Content   string    `json:"content"`
// 	CreatedAt time.Time `json:"created_at"`

// 	// User Users `gorm:"foreignKey:UserID"`
// }

// func main() {
// 	expected := uintptr(
// 		unsafe.Sizeof(uint(0)) + // ID
// 			unsafe.Sizeof(string("")) + // Content
// 			unsafe.Sizeof(uint(0)) + // UserID
// 			unsafe.Sizeof(uint(0)) + // PostID
// 			unsafe.Sizeof(time.Time{}), // CreatedAt
// 	)

// 	actual := unsafe.Sizeof(Comments{})

// 	fmt.Println("Expected:", expected, "bytes")
// 	fmt.Println("Actual:  ", actual, "bytes")

// 	fmt.Println("ID:", unsafe.Offsetof(Comments{}.ID))
// 	fmt.Println("UserID:", unsafe.Offsetof(Comments{}.UserID))
// 	fmt.Println("PostID:", unsafe.Offsetof(Comments{}.PostID))
// 	fmt.Println("Content:", unsafe.Offsetof(Comments{}.Content))
// 	fmt.Println("CreatedAt:", unsafe.Offsetof(Comments{}.CreatedAt))
// }
