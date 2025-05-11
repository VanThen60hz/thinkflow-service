package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

type rpcNoteShareLinkClient struct {
	client pb.NoteShareLinkServiceClient
}

func NewNoteShareLinkClient(client pb.NoteShareLinkServiceClient) *rpcNoteShareLinkClient {
	return &rpcNoteShareLinkClient{
		client: client,
	}
}

func (c *rpcNoteShareLinkClient) AddNewNoteShareLink(ctx context.Context, data *pb.NoteShareLinkCreation) error {
	req := &pb.CreateNoteShareLinkRequest{
		ShareLink: data,
	}

	if _, err := c.client.CreateNoteShareLink(ctx, req); err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	return nil
}

func (c *rpcNoteShareLinkClient) GetNoteShareLinkByID(ctx context.Context, id int64) (*pb.NoteShareLink, error) {
	req := &pb.GetNoteShareLinkByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetNoteShareLinkByID(ctx, req)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return resp.ShareLink, nil
}

func (c *rpcNoteShareLinkClient) GetNoteShareLinkByToken(ctx context.Context, token string) (*pb.NoteShareLink, error) {
	req := &pb.GetNoteShareLinkByTokenRequest{
		Token: token,
	}

	resp, err := c.client.GetNoteShareLinkByToken(ctx, req)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return resp.ShareLink, nil
}

func (c *rpcNoteShareLinkClient) UpdateNoteShareLink(ctx context.Context, id int64, data *pb.NoteShareLinkUpdate) error {
	req := &pb.UpdateNoteShareLinkRequest{
		Id:        id,
		ShareLink: data,
	}

	if _, err := c.client.UpdateNoteShareLink(ctx, req); err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	return nil
}

func (c *rpcNoteShareLinkClient) DeleteNoteShareLink(ctx context.Context, id int64) error {
	req := &pb.DeleteNoteShareLinkRequest{
		Id: id,
	}

	resp, err := c.client.DeleteNoteShareLink(ctx, req)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if !resp.Success {
		return core.ErrBadRequest.WithError("cannot delete note share link")
	}

	return nil
}
