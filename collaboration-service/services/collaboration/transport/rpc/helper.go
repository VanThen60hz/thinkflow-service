package rpc

import (
	"thinkflow-service/proto/pb"
	"thinkflow-service/services/collaboration/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPBCollaboration(collab *entity.Collaboration) *pb.Collaboration {
	return &pb.Collaboration{
		Id:         int32(collab.Id),
		NoteId:     int32(collab.NoteId),
		UserId:     int32(collab.UserId),
		Permission: string(collab.Permission),
		CreatedAt:  timestamppb.New(*collab.CreatedAt),
		UpdatedAt:  timestamppb.New(*collab.UpdatedAt),
	}
}

func toEntityPermissionType(permission string) entity.PermissionType {
	return entity.PermissionType(permission)
}
