package biz

import (
	"context"

	v1 "short-jump/api/user/v1"
	"short-jump/app/user/service/internal/conf"
)

type AuthUsecase struct {
	Key      string
	userRepo UserRepo
}

func NewAuthUsecase(conf *conf.Auth, userRepo UserRepo) *AuthUsecase {
	return &AuthUsecase{
		Key:      conf.ServiceKey,
		userRepo: userRepo,
	}
}

func (a *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// get user
	user, err := a.userRepo.FindByUsername(ctx, req.Email)
	if err != nil {
		// todo 处理error
		return nil, err
	}
}
