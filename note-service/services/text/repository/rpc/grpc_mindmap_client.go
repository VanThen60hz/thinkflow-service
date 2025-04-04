package rpc

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
	"gorm.io/datatypes"
)

type rpcMindmapClient struct {
	mindmapClient pb.MindmapServiceClient
}

func NewMindmapClient(client pb.MindmapServiceClient) *rpcMindmapClient {
	return &rpcMindmapClient{
		mindmapClient: client,
	}
}

func (c *rpcMindmapClient) GetMindmapById(ctx context.Context, id int64) (*common.SimpleMindmap, error) {
	resp, err := c.mindmapClient.GetMindmapById(ctx, &pb.GetMindmapByIdReq{Id: int32(id)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	mindMapDataJSON := datatypes.JSON(resp.Mindmap.MindmapData)

	mindmap := common.NewSimpleMindmap(
		int(resp.Mindmap.Id),
		mindMapDataJSON,
	)

	return &mindmap, nil
}
