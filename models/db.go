// models/db.go

package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:REPLACEPASSWORD@tcp(localhost:3306)/forumplsda?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Auto Migrate
	db.AutoMigrate(&User{}, &Thread{}, &Comment{}, &Category{})

	DB = db
}
