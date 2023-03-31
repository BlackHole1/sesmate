package model

import (
	"time"
)

type Version struct {
	ID        uint      `gorm:"primaryKey"`
	Version   string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"->,autoCreateTime"`
}
