package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

type rpcImageClient struct {
	imageClient pb.ImageServiceClient
}

func NewClient(imgClient pb.ImageServiceClient) *rpcImageClient {
	return &rpcImageClient{imageClient: imgClient}
}

func (c *rpcImageClient) GetImageById(ctx context.Context, id int) (*core.Image, error) {
	resp, err := c.imageClient.GetImageById(ctx, &pb.GetImageByIdReq{Id: int32(id)})
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
