package entity

import "gorm.io/gorm"

type UserSegmentPair struct {
	ID        uint `gorm:"primary key;autoIncrement" json:"id"`
	UserID    uint `json:"user_id"`
	SegmentID uint `json:"segment_id"`
}

func MigrateUserSegmentPairs(db *gorm.DB) error {
	err := db.AutoMigrate(&UserSegmentPair{})
	return err
}
