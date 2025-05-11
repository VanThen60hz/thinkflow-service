package business

import (
	"context"
	"fmt"

	"thinkflow-service/proto/pb"
)

func (biz *business) sendNotificationToAudioMembers(ctx context.Context, note *pb.GetNoteByIdResp, requesterId int, notiType string, message string) {
	err := biz.notiRepo.CreateNotification(ctx, notiType, int64(requesterId), int64(note.UserId), message, nil)
	if err != nil {
		fmt.Printf("Failed to send notification to owner: %v\n", err)
	}

	collabs, err := biz.collabRepo.GetCollaborationByNoteId(ctx, int(note.Id), nil)
	if err != nil {
		fmt.Printf("Failed to get collaborators: %v\n", err)
	} else if collabs != nil {
		for _, collab := range collabs {
			if collab != nil {
				err = biz.notiRepo.CreateNotification(ctx, notiType, int64(requesterId), int64(collab.UserId), message, nil)
				if err != nil {
					fmt.Printf("Failed to send notification to collaborator %d: %v\n", collab.UserId, err)
				}
			}
		}
	}
}
