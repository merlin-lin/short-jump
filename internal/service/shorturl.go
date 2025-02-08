package service

import (
	"context"

	pb "short-jump/api/shorturl/v1"
)

type ShorturlService struct {
	pb.UnimplementedShorturlServer
}

func NewShorturlService() *ShorturlService {
	return &ShorturlService{}
}

func (s *ShorturlService) CreateShorturl(ctx context.Context, req *pb.CreateShorturlRequest) (*pb.CreateShorturlReply, error) {
	return &pb.CreateShorturlReply{}, nil
}
func (s *ShorturlService) UpdateShorturl(ctx context.Context, req *pb.UpdateShorturlRequest) (*pb.UpdateShorturlReply, error) {
	return &pb.UpdateShorturlReply{}, nil
}
func (s *ShorturlService) DeleteShorturl(ctx context.Context, req *pb.DeleteShorturlRequest) (*pb.DeleteShorturlReply, error) {
	return &pb.DeleteShorturlReply{}, nil
}
func (s *ShorturlService) GetShorturl(ctx context.Context, req *pb.GetShorturlRequest) (*pb.GetShorturlReply, error) {
	return &pb.GetShorturlReply{}, nil
}
func (s *ShorturlService) ListShorturl(ctx context.Context, req *pb.ListShorturlRequest) (*pb.ListShorturlReply, error) {
	return &pb.ListShorturlReply{}, nil
}
func (s *ShorturlService) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}
