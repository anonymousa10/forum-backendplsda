// models/user.go

package models

// User represents the model for a forum user
type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "users"
}
