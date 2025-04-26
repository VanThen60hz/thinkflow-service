package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) GetUsersByIds(ctx context.Context, req *pb.GetUsersByIdsReq) (*pb.PublicUsersInfoResp, error) {
	userIDs := make([]int, len(req.Ids))

	for i := range userIDs {
		userIDs[i] = int(req.Ids[i])
	}

	users, err := s.business.GetUsersByIds(ctx, userIDs)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	publicUserInfo := make([]*pb.PublicUserInfo, len(users))

	for i := range users {
		publicUserInfo[i] = &pb.PublicUserInfo{
			Id:        int32(users[i].Id),
			FirstName: users[i].FirstName,
			LastName:  users[i].LastName,
			Email:     users[i].Email,
		}
	}

	return &pb.PublicUsersInfoResp{Users: publicUserInfo}, nil
}
