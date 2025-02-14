package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type User struct {
	Id           int
	Username     string
	Email        string
	Phone        string
	AvatarUrl    string
	Bio          string
	Status       int8
	PasswordHash string
}

type UserLogin struct {
}

type UserRepo interface {
	GetUser(ctx context.Context, id int32) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error

	//VerifyPassword(ctx context.Context, u *User, password string) error
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
