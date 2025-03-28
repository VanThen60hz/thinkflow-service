package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

type rpcClient struct {
	client pb.ImageServiceClient
}

func NewClient(client pb.ImageServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) GetImageById(ctx context.Context, id int) (*core.Image, error) {
	resp, err := c.client.GetImageById(ctx, &pb.GetImageByIdReq{Id: int32(id)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	image := core.NewImage(
		int(resp.Image.Id),
		resp.Image.Url,
		resp.Image.Width,
		resp.Image.Height,
		resp.Image.Extension,
	)

	return image, nil
}
