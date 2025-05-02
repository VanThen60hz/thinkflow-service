package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

type rpcAudioClient struct {
	audioClient pb.AudioServiceClient
}

func NewAudioClient(client pb.AudioServiceClient) *rpcAudioClient {
	return &rpcAudioClient{
		audioClient: client,
	}
}

func (c *rpcAudioClient) GetAudioById(ctx context.Context, id int64) (*pb.PublicAudioInfo, error) {
	resp, err := c.audioClient.GetAudioById(ctx, &pb.GetAudioByIdReq{Id: id})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return resp.Audio, nil
}

func (c *rpcAudioClient) GetAudiosByNoteId(ctx context.Context, noteId int64) ([]*pb.PublicAudioInfo, error) {
	resp, err := c.audioClient.GetAudiosByNoteId(ctx, &pb.GetAudiosByNoteIdReq{NoteId: noteId})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp == nil || resp.Audios == nil {
		return []*pb.PublicAudioInfo{}, nil
	}

	return resp.Audios, nil
}

func (c *rpcAudioClient) DeleteAudio(ctx context.Context, id int64) error {
	_, err := c.audioClient.DeleteAudio(ctx, &pb.DeleteAudioReq{Id: id})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
