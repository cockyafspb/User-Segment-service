package entity

import "gorm.io/gorm"

type Segment struct {
	ID   uint    `gorm:"primary key;autoIncrement" json:"id"`
	Slug *string `json:"slug"`
}

func MigrateSegments(db *gorm.DB) error {
	err := db.AutoMigrate(&Segment{})
	return err
}
