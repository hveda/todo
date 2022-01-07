package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        uint16           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title string `gorm:"type:varchar(100)" json:"title"`
	Email string `gorm:"type:varchar(100)" json:"email"`
}

type Todo struct {
	ID        uint16           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ActivityGroupId uint16 `gorm:"not null" json:"activity_group_id"`
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	IsActive bool `gorm:"type:boolean;default:true" json:"is_active"`
	Priority string `gorm:"type:varchar(16);default:very-high" json:"priority"`
}