package rpc

import (
	"context"
	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (c *rpcClient) DeleteUser(ctx context.Context, id int) error {
	resp, err := c.client.DeleteUser(ctx, &pb.DeleteUserReq{
		Id: int32(id),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	if !resp.Success {
		return errors.New("failed to delete user")
	}

	return nil
}
