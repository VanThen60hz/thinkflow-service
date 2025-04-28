package rpc

import (
	"thinkflow-service/proto/pb"
	"thinkflow-service/services/note-share-links/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPBNoteShareLink(entity *entity.NoteShareLink) *pb.NoteShareLink {
	if entity == nil {
		return nil
	}

	pbLink := &pb.NoteShareLink{
		Id:         int64(entity.Id),
		NoteId:     int32(entity.NoteID),
		Token:      entity.Token,
		Permission: entity.Permission,
	}

	if entity.ExpiresAt != nil {
		pbLink.ExpiresAt = timestamppb.New(*entity.ExpiresAt)
	}
	if entity.CreatedAt != nil {
		pbLink.CreatedAt = timestamppb.New(*entity.CreatedAt)
	}
	if entity.UpdatedAt != nil {
		pbLink.UpdatedAt = timestamppb.New(*entity.UpdatedAt)
	}

	return pbLink
}

func toEntityNoteShareLinkCreation(pbLink *pb.NoteShareLinkCreation) *entity.NoteShareLinkCreation {
	creation := &entity.NoteShareLinkCreation{
		NoteID:     int64(pbLink.NoteId),
		Token:      pbLink.Token,
		Permission: pbLink.Permission,
	}
	if pbLink.ExpiresAt != nil {
		expiresAt := pbLink.ExpiresAt.AsTime()
		creation.ExpiresAt = &expiresAt
	}
	return creation
}

func toEntityNoteShareLinkUpdate(pbLink *pb.NoteShareLinkUpdate) *entity.NoteShareLinkUpdate {
	update := &entity.NoteShareLinkUpdate{}
	if pbLink.Permission != "" {
		update.Permission = &pbLink.Permission
	}
	if pbLink.ExpiresAt != nil {
		expiresAt := pbLink.ExpiresAt.AsTime()
		update.ExpiresAt = &expiresAt
	}
	return update
}
