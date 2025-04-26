package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetNoteById(ctx context.Context, noteId int) (*entity.Note, error) {
	data, err := biz.noteRepo.GetNoteById(ctx, noteId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
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
		return nil, core.ErrInternalServerError.
			WithError("invalid requester type in context")
	}

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if requesterId == data.UserId {
		data.Permission = "owner"
	} else {
		hasWritePermission, err := biz.collabRepo.HasWritePermission(ctx, noteId, requesterId)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetNote.Error()).
				WithDebug(err.Error())
		}

		hasReadPermission, err := biz.collabRepo.HasReadPermission(ctx, noteId, requesterId)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetNote.Error()).
				WithDebug(err.Error())
		}

		if hasWritePermission {
			data.Permission = "write"
		} else if hasReadPermission {
			data.Permission = "read"
		} else {
			return nil, core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwnerOrCollaborator.Error())
		}
	}

	user, err := biz.userRepo.GetUserById(ctx, data.UserId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	data.User = user

	if data.SummaryID != nil {
		summary, err := biz.summaryRepo.GetSummaryById(ctx, *data.SummaryID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetNote.Error()).
				WithDebug(err.Error())
		}
		data.Summary = summary
	}

	if data.MindmapID != nil {
		mindmap, err := biz.mindmapRepo.GetMindmapById(ctx, *data.MindmapID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetNote.Error()).
				WithDebug(err.Error())
		}
		data.Mindmap = mindmap
	}

	return data, nil
}
