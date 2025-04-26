package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) GetUserIdByEmail(ctx context.Context, req *pb.GetUserIdByEmailReq) (*pb.GetUserIdByEmailResp, error) {
	userId, err := s.business.GetUserIdByEmail(ctx, req.Email)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.GetUserIdByEmailResp{Id: int32(userId)}, nil
}
