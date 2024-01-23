// models/comment.go

package models

import (
	"time"
)

// Comment represents the model for a forum comment
type Comment struct {
	ID        uint      `json:"id" gorm:"autoIncrement"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	ThreadID  uint      `json:"thread_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName sets the table name for the Comment model
func (Comment) TableName() string {
	return "comments"
}
