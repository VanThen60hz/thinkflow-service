package rpc

import (
	"context"

	"thinkflow-service/proto/pb"
)

func (s *grpcService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	err := s.business.DeleteUser(ctx, int(req.Id))
	if err != nil {
		return &pb.DeleteUserResp{Success: false}, err
	}

	return &pb.DeleteUserResp{Success: true}, nil
}
