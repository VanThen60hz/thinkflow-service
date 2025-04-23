package rpc

import (
	"context"
	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

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