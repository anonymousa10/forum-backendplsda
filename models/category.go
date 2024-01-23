// models/category.go

package models

// Category represents the model for a forum category
type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// TableName sets the table name for the Category model
func (Category) TableName() string {
	return "categories"
}
