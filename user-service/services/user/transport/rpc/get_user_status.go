package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) GetUserStatus(ctx context.Context, req *pb.GetUserStatusReq) (*pb.GetUserStatusResp, error) {
	status, err := s.business.GetUserStatus(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.GetUserStatusResp{Status: status}, nil
}
