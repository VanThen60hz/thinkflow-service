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

func NewImageClient(imgClient pb.ImageServiceClient) *rpcImageClient {
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

func (c *rpcImageClient) DeleteImage(ctx context.Context, id int) error {
	resp, err := c.imageClient.DeleteImage(ctx, &pb.DeleteImageReq{Id: int32(id)})
	if err != nil {
		return errors.WithStack(err)
	}

	if !resp.Success {
		return core.ErrInternalServerError.WithError("Failed to delete image")
	}

	return nil
}
