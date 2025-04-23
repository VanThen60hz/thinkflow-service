package rpc

import (
	"context"
	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (c *rpcClient) GetUserIdByEmail(ctx context.Context, email string) (int, error) {
	resp, err := c.client.GetUserIdByEmail(ctx, &pb.GetUserIdByEmailReq{
		Email: email,
	})
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(resp.Id), nil
}