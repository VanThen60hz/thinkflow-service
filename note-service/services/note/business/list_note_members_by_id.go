package business

import (
	"context"
	"fmt"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListNoteMembersById(ctx context.Context, id int, paging *core.Paging) ([]entity.NoteMember, error) {
	note, err := biz.noteRepo.GetNoteById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.WithDebug(err.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	if paging != nil {
		paging.Process()
	}

	collaborations, err := biz.collabRepo.GetCollaborationByNoteId(ctx, id, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("Cannot list note members").
			WithDebug(err.Error())
	}

	userIds := make([]int, 0, len(collaborations))
	for _, collab := range collaborations {
		userIds = append(userIds, collab.UserId)
	}

	users := []core.SimpleUser{}
	if len(userIds) > 0 {
		users, err = biz.userRepo.GetUsersByIds(ctx, userIds)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError("Cannot fetch users for note members").
				WithDebug(err.Error())
		}
		for i := range users {
			users[i].Mask(common.MaskTypeUser)
		}
	}

	owner, err := biz.userRepo.GetUserById(ctx, note.UserId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("Cannot fetch owner of note").
			WithDebug(err.Error())
	}
	owner.Mask(common.MaskTypeUser)

	var result []entity.NoteMember

	result = append(result, entity.NoteMember{
		SimpleUser: owner,
		Role:       "owner",
		Permission: "all",
	})

	roleMap := make(map[int]string)
	for _, collab := range collaborations {
		permission := "read"
		if collab.Permission != "" {
			permission = string(collab.Permission)
		}
		roleMap[collab.UserId] = permission
	}

	for _, u := range users {
		result = append(result, entity.NoteMember{
			SimpleUser: &u,
			Role:       "collaborator",
			Permission: roleMap[u.Id],
		})
	}

	// Debug log
	fmt.Println("Fetched collaborations:", collaborations)
	fmt.Println("Fetched users:", users)
	fmt.Println("Note owner:", owner)

	return result, nil
}
