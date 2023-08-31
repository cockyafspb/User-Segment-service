package entity

import "gorm.io/gorm"

type User struct {
	ID uint `gorm:"primary key;autoIncrement" json:"id"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
