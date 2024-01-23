// models/thread.go

package models

import (
	"time"
)

// Thread represents the model for a forum thread
type Thread struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName sets the table name for the Thread model
func (Thread) TableName() string {
	return "threads"
}
