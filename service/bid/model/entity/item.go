package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemId   uint
	ItemName string
	TypeId   uint
	Type     Type
	BidId    uuid.UUID
	Bid      Bid
}
