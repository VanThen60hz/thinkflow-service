package rpc

import (
	"context"
	"database/sql"

	pb "thinkflow-service/proto/pb"
	"thinkflow-service/services/notification/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Business interface {
	CreateNotification(ctx context.Context, data *entity.NotificationCreation) error
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	data := &entity.NotificationCreation{
		NotiType:       req.NotiType,
		NotiSenderID:   req.NotiSenderId,
		NotiReceivedID: req.NotiReceivedId,
		NotiContent:    req.NotiContent,
	}

	if req.NotiOptions != nil {
		data.NotiOptions = sql.NullString{
			String: *req.NotiOptions,
			Valid:  true,
		}
	}

	if err := s.business.CreateNotification(ctx, data); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateNotificationResponse{
		Id:             data.FakeId.String(),
		NotiType:       data.NotiType,
		NotiSenderId:   data.NotiSenderID,
		NotiReceivedId: data.NotiReceivedID,
		NotiContent:    data.NotiContent,
		NotiOptions:    &data.NotiOptions.String,
		IsRead:         data.IsRead,
		CreatedAt:      data.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      data.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
