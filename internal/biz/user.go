package biz

// 1、定义model对象
//2、定义查询account相关信息的repository
//3、为repository编写查询相关信息的抽象方法
//4、实例化repository
//5、声明可用于自动注入的UseCase
//6、编写调用respository(dao)层实现的方法

import (
	"context"

	pb "short-jump/api/user/v1"
)

/*
*
定义account对象
*/
// UsersInfo
type User struct {
	Id           int32  `json:"id"`            // 用户ID
	Email        string `json:"email"`         // 用户邮箱
	Phone        string `json:"phone"`         // 用户手机号（可选）
	Username     string `json:"username"`      // 用户名
	PasswordHash string `json:"password_hash"` // 密码的哈希值
	AvatarUrl    string `json:"avatar_url"`    // 用户头像URL
	Bio          string `json:"bio"`           // 用户简介（可选）
	CreatedAt    string `json:"created_at"`    // 注册时间
	UpdatedAt    string `json:"updated_at"`    // 最后更新时间
	LastLoginAt  string `json:"last_login_at"` // 上次登录时间
	Status       int8   `json:"status"`        // 账户状态，1 表示正常，0 表示禁用
	DeletedAt    string `json:"deleted_at"`    // 软删除字段，非 NULL 表示已删除
}

type UserRepo interface {
	FindById(ctx context.Context, loginReq *pb.GetUserRequest) (acc *User, err error)
}

type UserUsecase struct {
	uuc UserRepo
}

func NewUserUsecase(ucc UserRepo) *UserUsecase {
	return &UserUsecase{ucc}
}

func (uc *UserUsecase) GetAccountInfoById(ctx context.Context, getUserReq *pb.GetUserRequest) (acc *User, err error) {
	res, err := uc.uuc.FindById(ctx, getUserReq)
	return res, err
}
