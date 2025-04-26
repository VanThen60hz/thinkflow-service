package rpc

import (
	"context"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.NewUserIdResp, error) {
	newUserData := entity.NewUserForCreation(
		req.FirstName,
		req.LastName,
		req.Email,
	)

	if err := s.business.CreateNewUser(ctx, &newUserData); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.NewUserIdResp{Id: int32(newUserData.Id)}, nil
}
