package biz

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	v1 "short-jump/api/user/v1"
	"short-jump/app/user/service/internal/conf"
)

type AuthUsecase struct {
	key      string
	userRepo UserRepo
}

func NewAuthUsecase(userRepo UserRepo, conf *conf.Auth) *AuthUsecase {
	return &AuthUsecase{
		key:      conf.ApiKey,
		userRepo: userRepo,
	}
}

func (a *AuthUsecase) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.LoginResponse, error) {
	//  验证
	_, err := a.userRepo.FindByEmail(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: hashPassword(req.Password),
		Phone:        req.Phone,
		AvatarUrl:    req.AvatarUrl,
		Bio:          req.Bio,
	}
	// create user
	if err := a.userRepo.CreateUser(ctx, user); err != nil {
		return &v1.LoginResponse{Token: "lajjj"}, err
	}

	// create token
	token, err := genJWTToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &v1.LoginResponse{
		Token: token,
		User: &v1.GetUserResponse{
			Id:        0,
			Email:     user.Email,
			Phone:     user.Phone,
			Username:  user.Username,
			AvatarUrl: user.AvatarUrl,
			Bio:       user.Bio,
		},
	}, nil
}
func (a *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// get user
	user, err := a.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		// todo 处理error
		return nil, err
	}

	if err := verifyPassword(req.Password, user.PasswordHash); err != nil {
		return nil, err
	}

	token, err := genJWTToken(user.Id)
	if err != nil {
		return nil, err
	}
	return &v1.LoginResponse{
		Token: token,
		User: &v1.GetUserResponse{
			Id:        0,
			Email:     user.Email,
			Phone:     user.Phone,
			Username:  user.Username,
			AvatarUrl: user.AvatarUrl,
			Bio:       user.Bio,
		},
	}, nil
}

func hashPassword(s string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(password)
}

func verifyPassword(pwd, pwdHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwd))
}

func genJWTToken(userId int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})

	signedString, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
