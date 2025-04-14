package business

import (
	"context"
	"fmt"
	"time"

	"thinkflow-service/common"
	noteShareLinkEntity "thinkflow-service/services/note-share-links/entity"
	noteEntity "thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/google/uuid"
)

func (biz *business) CreateNoteShareLink(ctx context.Context, noteId int64, permission string, expiresAt *time.Time) (*noteShareLinkEntity.NoteShareLink, error) {
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

	noteData, err := biz.noteRepo.GetNoteById(ctx, int(noteId))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.WithError(noteEntity.ErrNoteNotFound.Error())
		}
		return nil, core.ErrInternalServerError.WithError(noteEntity.ErrCannotGetNote.Error()).WithDebug(err.Error())
	}

	if requesterId != noteData.UserId {
		return nil, core.ErrForbidden.WithError(noteEntity.ErrRequesterCannotModify.Error())
	}

	tid := uuid.New().String()
	uid = core.NewUID(uint32(noteId), 1, 1)
	sub := uid.String()

	token, _, err := biz.jwtProvider.IssueToken(ctx, tid, sub)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(noteShareLinkEntity.ErrCannotCreateShareLink.Error()).WithDebug(err.Error())
	}

	data := &noteShareLinkEntity.NoteShareLinkCreation{
		NoteID:     noteId,
		Permission: permission,
		Token:      token,
		ExpiresAt:  expiresAt,
	}

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return nil, core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.noteShareLinkRepo.AddNewNoteShareLink(ctx, data); err != nil {
		return nil, core.ErrInternalServerError.WithError(noteShareLinkEntity.ErrCannotCreateShareLink.Error()).WithDebug(err.Error())
	}

	key := fmt.Sprintf("share:%s", token)
	value := fmt.Sprintf("%d|%s", noteId, permission)

	ttl := time.Hour * 24
	if expiresAt != nil {
		expirySeconds := int(time.Until(*expiresAt).Seconds())
		if expirySeconds > 0 {
			ttl = min(ttl, time.Duration(expirySeconds)*time.Second)
		}
	}

	if err := biz.redisClient.Set(ctx, key, value, ttl); err != nil {
		return nil, core.ErrInternalServerError.WithError(noteShareLinkEntity.ErrCannotCreateShareLink.Error()).WithDebug(err.Error())
	}

	link := &noteShareLinkEntity.NoteShareLink{
		SQLModel:   data.SQLModel,
		NoteID:     data.NoteID,
		Permission: data.Permission,
		Token:      data.Token,
		ExpiresAt:  data.ExpiresAt,
		CreatedBy:  data.CreatedBy,
	}

	return link, nil
}
