package rpc

import (
	"context"
	"log"

	"thinkflow-service/proto/pb"
)

func (s *grpcService) DeleteAuth(ctx context.Context, req *pb.DeleteAuthReq) (*pb.DeleteAuthResp, error) {
	err := s.business.DeleteAuth(ctx, int(req.UserId))
	if err != nil {
		log.Printf("DeleteAuth error: %v", err)
		return &pb.DeleteAuthResp{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.DeleteAuthResp{
		Success: true,
	}, nil
}
