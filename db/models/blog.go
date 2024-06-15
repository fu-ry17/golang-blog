package models

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
