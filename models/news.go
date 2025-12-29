package models

import "time"

type News struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"`
}
