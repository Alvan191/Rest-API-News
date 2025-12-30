package models

import "time"

type News struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Title     string     `json:"title" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	Author    string     `json:"author"` //data yang di migrate
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `gorm:"<-:update" json:"updated_at"`
}
