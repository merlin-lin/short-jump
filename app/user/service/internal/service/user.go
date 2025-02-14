package service

import (
	"context"

	pb "short-jump/api/user/v1"
)

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.ac.Login(ctx, req)
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.LoginResponse, error) {
	return s.ac.Register(ctx, req)
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	x, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:        int32(x.Id),
		Username:  x.Username,
		Email:     x.Email,
		Phone:     x.Phone,
		AvatarUrl: x.AvatarUrl,
		Bio:       x.Bio,
		Status:    int32(x.Status),
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}
func (s *UserService) DisableUser(ctx context.Context, req *pb.DisableUserRequest) (*pb.DisableUserResponse, error) {
	return &pb.DisableUserResponse{}, nil
}
func (s *UserService) EnableUser(ctx context.Context, req *pb.EnableUserRequest) (*pb.EnableUserResponse, error) {
	return &pb.EnableUserResponse{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}
