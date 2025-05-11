package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

type grpcCollaborationClient struct {
	client pb.CollaborationServiceClient
}

func NewCollaborationClient(client pb.CollaborationServiceClient) *grpcCollaborationClient {
	return &grpcCollaborationClient{
		client: client,
	}
}

func (c *grpcCollaborationClient) AddNewCollaboration(ctx context.Context, data *pb.CollaborationCreation) error {
	req := &pb.AddCollaborationRequest{
		Collaboration: data,
	}

	resp, err := c.client.AddCollaboration(ctx, req)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if !resp.Success {
		return core.ErrBadRequest.WithError("cannot create collaboration")
	}

	return nil
}

func (c *grpcCollaborationClient) HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error) {
	req := &pb.CheckReadPermissionRequest{
		NoteId: int32(noteId),
		UserId: int32(userId),
	}

	resp, err := c.client.CheckReadPermission(ctx, req)
	if err != nil {
		return false, core.ErrInternalServerError.WithError(err.Error())
	}

	return resp.HasPermission, nil
}

func (c *grpcCollaborationClient) HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error) {
	req := &pb.CheckWritePermissionRequest{
		NoteId: int32(noteId),
		UserId: int32(userId),
	}

	resp, err := c.client.CheckWritePermission(ctx, req)
	if err != nil {
		return false, core.ErrInternalServerError.WithError(err.Error())
	}

	return resp.HasPermission, nil
}

func (c *grpcCollaborationClient) GetCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) (*pb.Collaboration, error) {
	req := &pb.GetCollaborationByNoteIdAndUserIdRequest{
		NoteId: int32(noteId),
		UserId: int32(userId),
	}

	resp, err := c.client.GetCollaborationByNoteIdAndUserId(ctx, req)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	if resp.Collaboration == nil {
		return nil, core.ErrNotFound.WithError("collaboration not found")
	}

	return resp.Collaboration, nil
}

func (c *grpcCollaborationClient) GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]*pb.Collaboration, error) {
	req := &pb.GetCollaborationByNoteIdRequest{
		NoteId: int32(noteId),
	}

	if paging != nil {
		req.Page = int32(paging.Page)
		req.Limit = int32(paging.Limit)
	}

	resp, err := c.client.GetCollaborationByNoteId(ctx, req)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	if paging != nil {
		paging.Total = int64(resp.Total)
	}
	return resp.Collaborations, nil
}

func (c *grpcCollaborationClient) GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]*pb.Collaboration, error) {
	req := &pb.GetCollaborationByUserIdRequest{
		UserId: int32(userId),
		Page:   int32(paging.Page),
		Limit:  int32(paging.Limit),
	}

	resp, err := c.client.GetCollaborationByUserId(ctx, req)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	paging.Total = int64(resp.Total)
	return resp.Collaborations, nil
}

func (c *grpcCollaborationClient) UpdateCollaboration(ctx context.Context, id int, data *pb.CollaborationUpdate) error {
	req := &pb.UpdateCollaborationRequest{
		Id:            int32(id),
		Collaboration: data,
	}

	resp, err := c.client.UpdateCollaboration(ctx, req)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if !resp.Success {
		return core.ErrBadRequest.WithError("cannot update collaboration")
	}

	return nil
}

func (c *grpcCollaborationClient) DeleteCollaboration(ctx context.Context, id int) error {
	req := &pb.DeleteCollaborationRequest{
		Id: int32(id),
	}

	resp, err := c.client.DeleteCollaboration(ctx, req)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if !resp.Success {
		return core.ErrBadRequest.WithError("cannot delete collaboration")
	}

	return nil
}

func (c *grpcCollaborationClient) RemoveCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) error {
	req := &pb.RemoveCollaborationByNoteIdAndUserIdRequest{
		NoteId: int32(noteId),
		UserId: int32(userId),
	}

	resp, err := c.client.RemoveCollaborationByNoteIdAndUserId(ctx, req)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if !resp.Success {
		return core.ErrBadRequest.WithError("cannot remove collaboration")
	}

	return nil
}
