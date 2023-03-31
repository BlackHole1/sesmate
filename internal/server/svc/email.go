package svc

import (
	"gorm.io/gorm"

	"github.com/BlackHole1/sesmate/internal/server/model"
	"github.com/BlackHole1/sesmate/internal/server/sqlite"
)

type emailContext struct {
	db *gorm.DB
}

var Email = &emailContext{
	db: sqlite.Client(),
}

func (c *emailContext) Record(p *model.EmailRecord) error {
	return c.db.Create(p).Error
}

func (c *emailContext) AllRecord() ([]*model.EmailRecord, error) {
	var p []*model.EmailRecord
	err := c.db.Where("1=1").Find(&p).Error

	return p, err
}

func (c *emailContext) DeleteAllRecord() error {
	return c.db.Where("1=1").Delete(&model.EmailRecord{}).Error
}
