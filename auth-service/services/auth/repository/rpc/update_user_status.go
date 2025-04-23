package rpc

import (
	"context"
	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

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