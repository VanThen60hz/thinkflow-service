package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteNoteMember(ctx context.Context, noteId int, userId int) error {
	note, err := biz.noteRepo.GetNoteById(ctx, noteId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.WithDebug(err.Error())
		}
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
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
		return core.ErrForbidden.WithError("you have not permission to delete member this note (just only owner can modify team members)")
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
		return core.ErrBadRequest.WithError("Cannot delete owner from note")
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

	if err := biz.collabRepo.DeleteCollaboration(ctx, collab.Id); err != nil {
		return core.ErrInternalServerError.
			WithError("Cannot delete collaboration").
			WithDebug(err.Error())
	}

	return nil
}
