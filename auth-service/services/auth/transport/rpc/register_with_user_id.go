package rpc

import (
	"context"
	"log"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/auth/entity"
)

func (s *grpcService) RegisterWithUserId(ctx context.Context, req *pb.RegisterWithUserIdReq) (*pb.RegisterWithUserIdResp, error) {
	data := &entity.AuthRegister{
		AuthEmailPassword: entity.AuthEmailPassword{
			Email:    req.Email,
			Password: req.Password,
		},
	}

	err := s.business.RegisterWithUserId(ctx, data, int(req.UserId))
	if err != nil {
		log.Printf("RegisterWithUserId error: %v", err)
		return &pb.RegisterWithUserIdResp{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.RegisterWithUserIdResp{
		Success: true,
	}, nil
}
