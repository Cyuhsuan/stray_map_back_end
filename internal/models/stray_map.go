package models

import (
	"time"
)

type StrayMap struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Title       string    `json:"title" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Image       string    `json:"image"`
	Location    string    `json:"location" gorm:"type:varchar(200);not null"`
	Latitude    float64   `json:"latitude" gorm:"default:0.0"`
	Longitude   float64   `json:"longitude" gorm:"default:0.0"`
	PetType     string    `json:"pet_type" gorm:"type:varchar(20);default:other"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
