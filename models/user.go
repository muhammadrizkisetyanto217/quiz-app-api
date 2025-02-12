package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}