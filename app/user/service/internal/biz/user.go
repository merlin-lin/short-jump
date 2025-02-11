package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id        int
	Username  string
	Email     string
	Phone     string
	AvatarUrl string
	Bio       string
	Status    int8
}

type UserRepo interface {
	GetUser(ctx context.Context, id int32) (*User, error)
}

type UserUsecase struct {
	repo UserRepo

	log *log.Helper
}

func (uc *UserUsecase) Get(ctx context.Context, id int32) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}
