package model

import (
	"time"
)

type EmailRecord struct {
	RequestId string `gorm:"primaryKey"`
	Data      string
	RawData   string
	MessageId string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"->,autoCreateTime"`
}
