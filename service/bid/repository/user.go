package repository

import (
	"github.com/leehai1107/bipbip/service/bid/model/entity"
	"gorm.io/gorm"
)

type IUserRepo interface {
	GetUserByUserName(userName string) (*entity.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUserByUserName(userName string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
