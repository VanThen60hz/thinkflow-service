package rpc

import (
	"context"

	"thinkflow-service/proto/pb"
)

type rpcNoteClient struct {
	noteClient pb.NoteServiceClient
}

func NewNoteClient(noteClient pb.NoteServiceClient) *rpcNoteClient {
	return &rpcNoteClient{noteClient: noteClient}
}

func (client *rpcNoteClient) DeleteUserNotes(ctx context.Context, userId int32) (bool, int32, error) {
	req := &pb.DeleteUserNotesReq{
		UserId: userId,
	}

	resp, err := client.noteClient.DeleteUserNotes(ctx, req)
	if err != nil {
		return false, 0, err
	}

	return resp.Success, resp.DeletedCount, nil
}
