package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	pb "short-jump/api/user/v1"
	"short-jump/app/user/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUsecase
	ac  *biz.AuthUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, ac *biz.AuthUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/user")),
	}
}
