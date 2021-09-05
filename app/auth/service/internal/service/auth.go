package service

import (
	"context"
	"scrm/app/auth/service/internal/biz"

	pb "scrm/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	ac *biz.AuthUseCase
}

func NewAuthService(ac *biz.AuthUseCase) *AuthService {
	return &AuthService{
		ac: ac,
	}
}

func (s *AuthService) GetToken(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	a, _ := s.ac.GetToken(ctx, req.Username, req.Password)
	return &pb.LoginReply{
		Token: a.Token,
	}, nil
}
