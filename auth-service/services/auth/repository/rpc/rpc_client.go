package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type rpcClient struct {
	client pb.UserServiceClient
	db     *gorm.DB
}

func NewClient(client pb.UserServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error) {
	resp, err := c.client.CreateUser(ctx, &pb.CreateUserReq{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	})
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(resp.Id), nil
}

func (c *rpcClient) GetDB() *gorm.DB {
	return c.db
}

func (c *rpcClient) SetDB(db *gorm.DB) {
	c.db = db
}

func (c *rpcClient) GetUserIdByEmail(ctx context.Context, email string) (int, error) {
	// TODO: Implement this when user service supports it
	return 0, errors.New("not implemented")
}

func (c *rpcClient) UpdateUserStatus(ctx context.Context, id int, status string) error {
	_, err := c.client.UpdateUserStatus(ctx, &pb.UpdateUserStatusReq{
		Id:     int32(id),
		Status: status,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *rpcClient) GetUserStatus(ctx context.Context, id int) (string, error) {
	resp, err := c.client.GetUserStatus(ctx, &pb.GetUserStatusReq{
		Id: int32(id),
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	return resp.Status, nil
}
