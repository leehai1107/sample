package usecase

import (
	"context"

	"github.com/leehai1107/bipbip/service/bid/model/request"
	"github.com/leehai1107/bipbip/service/bid/repository"
)

type IUserUsecase interface {
	Login(ctx context.Context, req request.Login) (string, error)
}

type userUsecase struct {
	repo repository.IUserRepo
}

func NewUserUsecase(
	repo repository.IUserRepo,
) IUserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Login(ctx context.Context, req request.Login) (string, error) {
	res := "Login success: "
	user, err := u.repo.GetUserByUserName(req.Username)
	if err != nil {
		return "Login failed: " + req.Username, err
	}
	return res + user.Username, nil
}
