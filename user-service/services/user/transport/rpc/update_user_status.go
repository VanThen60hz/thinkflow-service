package rpc

import (
	"context"

	"thinkflow-service/proto/pb"
)

func (s *grpcService) UpdateUserStatus(ctx context.Context, req *pb.UpdateUserStatusReq) (*pb.UpdateUserStatusResp, error) {
	err := s.business.UpdateUserStatus(ctx, int(req.Id), req.Status)
	if err != nil {
		return &pb.UpdateUserStatusResp{Success: false}, err
	}

	return &pb.UpdateUserStatusResp{Success: true}, nil
}
