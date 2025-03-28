package rpc

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

type rpcTranscriptClient struct {
	transcriptClient pb.TranscriptServiceClient
}

func NewTranscriptClient(client pb.TranscriptServiceClient) *rpcTranscriptClient {
	return &rpcTranscriptClient{
		transcriptClient: client,
	}
}

func (c *rpcTranscriptClient) GetTranscriptById(ctx context.Context, id int64) (*common.SimpleTranscript, error) {
	resp, err := c.transcriptClient.GetTranscriptById(ctx, &pb.GetTranscriptByIdReq{Id: int32(id)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	transcript := common.NewSimpleTranscript(
		int(resp.Transcript.Id),
		resp.Transcript.Content,
	)

	return &transcript, nil
}
