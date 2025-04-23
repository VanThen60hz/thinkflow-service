package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (c *rpcClient) GetUserStatus(ctx context.Context, id int) (string, error) {
	resp, err := c.client.GetUserStatus(ctx, &pb.GetUserStatusReq{
		Id: int32(id),
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	return resp.Status, nil
}
