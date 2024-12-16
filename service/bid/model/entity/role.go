package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleId   uint
	RoleName string
	UserId   uuid.UUID
	User     User
}
