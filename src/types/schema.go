package types

import (
	"time"

	// "gorm.io/gorm"
)

type Activity struct {
	ID int `gorm:"AUTO_INCREMENT;primary_key;index" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	Title string `json:"title"`
	Email string `json:"email"`
}

type ToDo struct {
	ID int `gorm:"AUTO_INCREMENT;primary_key;index" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	ActivityGroupId string `gorm:"not null" json:"activity_group_id"`
	Title string `gorm:"not null" json:"title"`
	IsActive bool `gorm:"default:true" json:"is_active"`
	Priority string `json:"priority"`
}