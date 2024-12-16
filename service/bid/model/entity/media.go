package entity

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	MediaId uint
	Url     string
	ItemId  uint
	Item    Item
}
