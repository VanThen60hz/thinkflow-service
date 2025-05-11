package business

import (
	"context"
	"fmt"

	noteEntity "thinkflow-service/services/note/entity"
)

func (biz *business) sendNotificationToNoteMembers(ctx context.Context, note *noteEntity.Note, requesterId int, notiType string, message string) {
	if int64(note.UserId) != int64(requesterId) {
		err := biz.notiRepo.CreateNotification(ctx, notiType, int64(requesterId), int64(note.UserId), message, nil)
		if err != nil {
			fmt.Printf("Failed to send notification to owner: %v\n", err)
		}
	}

	collabs, err := biz.collabRepo.GetCollaborationByNoteId(ctx, note.Id, nil)
	if err != nil {
		fmt.Printf("Failed to get collaborators: %v\n", err)
	} else if collabs != nil {
		for _, collab := range collabs {
			if collab != nil && collab.UserId != int32(requesterId) {
				err = biz.notiRepo.CreateNotification(ctx, notiType, int64(requesterId), int64(collab.UserId), message, nil)
				if err != nil {
					fmt.Printf("Failed to send notification to collaborator %d: %v\n", collab.UserId, err)
				}
			}
		}
	}
}
