package business

import (
	"fmt"	
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListNoteMembersByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]entity.NoteMember, error) {
	note, err := biz.noteRepo.GetNoteById(ctx, noteId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.WithDebug(err.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return nil, core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return nil, core.ErrInternalServerError.WithError("invalid requester type in context")
	}

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasPermissionRead, err := biz.collabRepo.HasReadPermission(ctx, noteId, requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	if requesterId != note.UserId && !hasPermissionRead {
		return nil, core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwnerOrCollaborator.Error())
	}

	if paging != nil {
		paging.Process()
	}

	collaborations, err := biz.collabRepo.GetCollaborationByNoteId(ctx, noteId, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("Cannot list note members").
			WithDebug(err.Error())
	}

	userIds := make([]int, 0, len(collaborations))
	for _, collab := range collaborations {
		userIds = append(userIds, int(collab.UserId))
	}

	users := []core.SimpleUser{}
	if len(userIds) > 0 {
		users, err = biz.userRepo.GetUsersByIds(ctx, userIds)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError("Cannot fetch users for note members").
				WithDebug(err.Error())
		}

		avatarIds := make([]int, 0)
		for i := range users {
			if users[i].AvatarId > 0 {
				avatarIds = append(avatarIds, users[i].AvatarId)
			}
			users[i].Mask(common.MaskTypeUser)
		}

		avatarMap := make(map[int]*core.Image)
		for _, avatarId := range avatarIds {
			img, err := biz.imageRepo.GetImageById(ctx, avatarId)
			if err != nil {
				return nil, core.ErrInternalServerError.
					WithError("cannot get user avatar").
					WithDebug(err.Error())
			}
			avatarMap[avatarId] = img
		}

		for i := range users {
			if users[i].AvatarId > 0 {
				if avatar, ok := avatarMap[users[i].AvatarId]; ok {
					users[i].Avatar = avatar
				}
			}
		}
	}

	owner, err := biz.userRepo.GetUserById(ctx, note.UserId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("Cannot fetch owner of note").
			WithDebug(err.Error())
	}

	fmt.Println("owner: ", owner)

	if owner.AvatarId > 0 {
		avatar, err := biz.imageRepo.GetImageById(ctx, owner.AvatarId)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError("cannot get user avatar").
				WithDebug(err.Error())
		}
		owner.Avatar = avatar
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
			permission = collab.Permission
		}
		roleMap[int(collab.UserId)] = permission
	}

	for _, u := range users {
		result = append(result, entity.NoteMember{
			SimpleUser: &u,
			Role:       "collaborator",
			Permission: roleMap[u.Id],
		})
	}

	return result, nil
}
