package sql

import (
	"github.com/leehai1107/bipbip/service/bid/model/entity"
	"gorm.io/gorm"
)

type IMigration interface {
	Migrate() error
}

type migration struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) IMigration {
	return &migration{db: db}
}

func (m *migration) Migrate() error {
	return m.db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.Media{},
		&entity.Type{},
		&entity.Item{},
		&entity.Bid{})
}
