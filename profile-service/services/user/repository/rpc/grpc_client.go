package rpc

import (
	"context"
	"fmt"

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

func (c *rpcClient) GetImagesByIds(ctx context.Context, ids []int) ([]core.Image, error) {
	ImageIds := make([]int32, len(ids))

	for i := range ids {
		ImageIds[i] = int32(ids[i])
	}

	fmt.Println(ImageIds, "ImageIds")

	resp, err := c.client.GetImagesByIds(ctx, &pb.GetImagesByIdsReq{Ids: ImageIds})

	fmt.Println("resp", resp)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	Images := make([]core.Image, len(resp.Images))

	for i := range Images {
		respImage := resp.Images[i]
		Images[i] = core.Image{
			Id:        int64(respImage.Id),
			Url:       respImage.Url,
			Width:     respImage.Width,
			Height:    respImage.Height,
			Extension: respImage.Extension,
		}
	}

	return Images, nil
}

func (c *rpcClient) GetImageById(ctx context.Context, id int) (*core.Image, error) {
	resp, err := c.client.GetImageById(ctx, &pb.GetImageByIdReq{Id: int32(id)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	Image := core.Image{
		Id:        int64(resp.Image.Id),
		Url:       resp.Image.Url,
		Width:     resp.Image.Width,
		Height:    resp.Image.Height,
		Extension: resp.Image.Extension,
	}
	return &Image, nil
}
