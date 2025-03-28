package rpc

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

type rpcSummaryClient struct {
	SummaryClient pb.SummaryServiceClient
}

func NewSummaryClient(client pb.SummaryServiceClient) *rpcSummaryClient {
	return &rpcSummaryClient{
		SummaryClient: client,
	}
}

func (c *rpcSummaryClient) GetSummaryById(ctx context.Context, id int64) (*common.SimpleSummary, error) {
	resp, err := c.SummaryClient.GetSummaryById(ctx, &pb.GetSummaryByIdReq{Id: int32(id)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	summary := common.NewSimpleSummary(
		int(resp.Summary.Id),
		resp.Summary.SummaryText,
	)

	return &summary, nil
}
