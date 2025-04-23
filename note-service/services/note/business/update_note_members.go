package business

import (
	"context"

	"thinkflow-service/common"
	collabEntity "thinkflow-service/services/collaboration/entity"
	noteEntity "thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateNoteMemberAccess(ctx context.Context, noteId int, userId int, permission string) error {
	note, err := biz.noteRepo.GetNoteById(ctx, noteId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.WithDebug(err.Error())
		}
		return core.ErrInternalServerError.
			WithError(noteEntity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return core.ErrInternalServerError.WithError("invalid requester type in context")
	}

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if requesterId != note.UserId {
		return core.ErrForbidden.WithError("you have not permission to edit member this note (just only owner can modify team members)")
	}

	_, err = biz.userRepo.GetUserById(ctx, userId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.WithError("User not found").WithDebug(err.Error())
		}
		return core.ErrInternalServerError.
			WithError("Cannot get user").
			WithDebug(err.Error())
	}

	if note.UserId == userId {
		return core.ErrBadRequest.WithError("Cannot edit owner's permission")
	}

	collab, err := biz.collabRepo.GetCollaborationByNoteIdAndUserId(ctx, noteId, userId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.WithError("User is not a member of this note").WithDebug(err.Error())
		}
		return core.ErrInternalServerError.
			WithError("Cannot get collaboration").
			WithDebug(err.Error())
	}

	collab.Permission = collabEntity.PermissionType(permission)
	if err := biz.collabRepo.UpdateCollaboration(ctx, collab.Id, collab); err != nil {
		return core.ErrInternalServerError.
			WithError("Cannot update collaboration").
			WithDebug(err.Error())
	}

	return nil
}
