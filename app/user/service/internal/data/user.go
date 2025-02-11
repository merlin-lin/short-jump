package data

import (
	context "context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"short-jump/app/user/service/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Status   int    `json:"status"`
}

func (u *userRepo) GetUser(ctx context.Context, id int32) (*biz.User, error) {
	user := User{}
	result := u.data.db.WithContext(ctx).Find(&user, id)

	return &biz.User{
		Id:        int(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		AvatarUrl: user.Avatar,
		Bio:       user.Bio,
		Status:    int8(user.Status),
	}, result.Error
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
