package user

import (
	"context"
	"fmt"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	GetUserDetails(ctx context.Context, id int) (*entity.User, error)
	GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error)
	CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error
	UpdateUserStatus(ctx context.Context, id int, status string) error
	GetUserStatus(ctx context.Context, id int) (string, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*pb.PublicUserInfoResp, error) {
	user, err := s.business.GetUserDetails(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.PublicUserInfoResp{
		User: &pb.PublicUserInfo{
			Id:        int32(user.Id),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}, nil
}

func (s *grpcService) GetUsersByIds(ctx context.Context, req *pb.GetUsersByIdsReq) (*pb.PublicUsersInfoResp, error) {
	userIDs := make([]int, len(req.Ids))

	for i := range userIDs {
		userIDs[i] = int(req.Ids[i])
	}

	fmt.Println("userIDs", userIDs)

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

func (s *grpcService) UpdateUserStatus(ctx context.Context, req *pb.UpdateUserStatusReq) (*pb.UpdateUserStatusResp, error) {
	err := s.business.UpdateUserStatus(ctx, int(req.Id), req.Status)
	if err != nil {
		return &pb.UpdateUserStatusResp{Success: false}, err
	}

	return &pb.UpdateUserStatusResp{Success: true}, nil
}

func (s *grpcService) GetUserStatus(ctx context.Context, req *pb.GetUserStatusReq) (*pb.GetUserStatusResp, error) {
	status, err := s.business.GetUserStatus(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.GetUserStatusResp{Status: status}, nil
}
