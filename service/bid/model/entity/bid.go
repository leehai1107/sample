package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	BidId        uuid.UUID
	BidName      string
	HostId       uint
	User         User `gorm:"foreignKey:HostId"`
	BinPrice     uint
	InitialPrice uint
	FinalPrice   uint
	StepPrice    uint
	StartTime    time.Time
	EndTime      time.Time
	Condition    string
	Description  string
	Guarantee    string
}
