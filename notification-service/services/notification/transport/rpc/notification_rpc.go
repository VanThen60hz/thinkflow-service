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
	CreateNotification(ctx context.Context, data *entity.NotificationCreation) (*entity.Notification, error)
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

	notification, err := s.business.CreateNotification(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateNotificationResponse{
		Id:             notification.FakeId.String(),
		NotiType:       notification.NotiType,
		NotiSenderId:   notification.NotiSenderID,
		NotiReceivedId: notification.NotiReceivedID,
		NotiContent:    notification.NotiContent,
		NotiOptions:    &notification.NotiOptions.String,
		IsRead:         notification.IsRead,
		CreatedAt:      notification.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      notification.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
