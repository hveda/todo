package types

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID int `gorm:"AUTO_INCREMENT;primary_key;index" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return nil
}

type Post struct {
	Base
	Title string `gorm:"not null" json:"title"`
	Slug string `gorm:"not null" json:"slug"`
	Content string `gorm:"type:text; not null" json:"content"`
	Summary string `gorm:"not null; default: ''" json:"summary"`
}

type Home struct {
	Base
	Message string `gorm:"type:text; not null" json:"message"`
}

type Activity struct {
	Base
	Title string `json:"title"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type ToDo struct {
	Base
	ActivityGroupId int64 `gorm:"not null" json:"activity_group_id"`
	Title string `gorm:"not null" json:"title"`
	IsActive bool `gorm:"default:true" json:"is_active"`
	Priority string `json:"priority"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}