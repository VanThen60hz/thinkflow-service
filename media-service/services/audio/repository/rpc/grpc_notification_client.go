package rpc

import (
	"context"

	pb "thinkflow-service/proto/pb"
)

type rpcNotificationClient struct {
	client pb.NotificationServiceClient
}

func NewNotificationClient(client pb.NotificationServiceClient) *rpcNotificationClient {
	return &rpcNotificationClient{
		client: client,
	}
}

func (c *rpcNotificationClient) CreateNotification(ctx context.Context, notiType string, senderId, receiverId int64, content string, options *string) error {
	req := &pb.CreateNotificationRequest{
		NotiType:       notiType,
		NotiSenderId:   senderId,
		NotiReceivedId: receiverId,
		NotiContent:    content,
	}

	if options != nil {
		req.NotiOptions = options
	}

	_, err := c.client.CreateNotification(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
