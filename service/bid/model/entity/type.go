package entity

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	TypeId   uint
	TypeName string
}
