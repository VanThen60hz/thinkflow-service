package rpc

import (
	"context"
	"fmt"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

type rpcAuthClient struct {
	authClient pb.AuthServiceClient
}

func NewAuthClient(authClient pb.AuthServiceClient) *rpcAuthClient {
	return &rpcAuthClient{authClient: authClient}
}

func (client *rpcAuthClient) RegisterWithUserId(ctx context.Context, userId int32, email, password string) error {
	req := &pb.RegisterWithUserIdReq{
		UserId:   userId,
		Email:    email,
		Password: password,
	}

	resp, err := client.authClient.RegisterWithUserId(ctx, req)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if !resp.Success {
		return fmt.Errorf("auth-service error: %s", resp.ErrorMessage)
	}

	return nil
}

func (client *rpcAuthClient) DeleteAuth(ctx context.Context, userId int32) error {
	req := &pb.DeleteAuthReq{
		UserId: userId,
	}

	resp, err := client.authClient.DeleteAuth(ctx, req)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if !resp.Success {
		return fmt.Errorf("auth-service error: %s", resp.ErrorMessage)
	}

	return nil
}
