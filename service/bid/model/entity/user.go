package entity

import (
	"github.com/google/uuid"
	"github.com/leehai1107/bipbip/pkg/utils/timeutils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId    uuid.UUID
	Username  string `gorm:"unique;not null"`
	Password  string
	Email     string
	FirstName string
	Dob       timeutils.Date
	Address   string
}
