package model

import (
	"time"

	"gorm.io/gorm"
)

type FCMToken struct {
	ID        uint   `gorm:"primarykey"`
	UserID    string `gorm:"index;not null"`
	Token     string `gorm:"uniqueIndex;not null"`
	DeviceID  string `gorm:"index;not null"`
	Platform  string `gorm:"not null"` // "android" or "ios"
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
