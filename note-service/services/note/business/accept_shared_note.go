package business

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"thinkflow-service/common"
	collabEntity "thinkflow-service/services/collaboration/entity"
	noteShareLinkEntity "thinkflow-service/services/note-share-links/entity"
	noteEntity "thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/redis/go-redis/v9"
)

func (biz *business) AcceptSharedNote(ctx context.Context, token string) (*noteEntity.Note, error) {
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

	noteId, permission, err := biz.getNoteShareData(ctx, token)
	if err != nil {
		return nil, err
	}

	isCollaboration, err := biz.collabRepo.HasReadPermission(ctx, noteId, requesterId)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound.WithError("collaboration not found")
		}
		return nil, core.ErrInternalServerError.WithError("cannot check collaboration").WithDebug(err.Error())
	}

	note, err := biz.noteRepo.GetNoteById(ctx, noteId)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound.WithError("note not found")
		}
		return nil, core.ErrInternalServerError.WithError("cannot get note").WithDebug(err.Error())
	}

	if note.UserId == requesterId || isCollaboration {
		return nil, core.ErrBadRequest.WithError("you already have access to this note")
	}

	newCollabData := &collabEntity.CollaborationCreation{
		UserId:     requesterId,
		NoteId:     noteId,
		Permission: collabEntity.PermissionType(permission),
	}

	if err := biz.collabRepo.AddNewCollaboration(ctx, newCollabData); err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot add collaboration").WithDebug(err.Error())
	}

	return note, nil
}

func (biz *business) getNoteShareData(ctx context.Context, token string) (int, string, error) {
	key := fmt.Sprintf("share:%s", token)
	storedValue, err := biz.redisClient.Get(ctx, key)
	if err == nil && storedValue != "" {
		parts := strings.Split(storedValue, "|")
		if len(parts) != 2 {
			return 0, "", core.ErrInternalServerError.WithError("invalid cache format")
		}
		noteId, convErr := strconv.Atoi(parts[0])
		if convErr != nil {
			return 0, "", core.ErrInternalServerError.WithError("cannot convert note id").WithDebug(convErr.Error())
		}
		return noteId, parts[1], nil
	}

	if !errors.Is(err, redis.Nil) {
		return 0, "", core.ErrInternalServerError.WithError("cannot get share link").WithDebug(err.Error())
	}

	noteSharedData, err := biz.noteShareLinkRepo.GetNoteShareLinkByToken(ctx, token)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return 0, "", core.ErrNotFound.WithError("share link not found")
		}

		if errors.Is(err, noteShareLinkEntity.ErrNoteShareLinkNotFound) {
			return 0, "", core.ErrNotFound.WithError("share link is deactivate or expired")
		}
		return 0, "", core.ErrInternalServerError.WithError("cannot get share link").WithDebug(err.Error())
	}

	if noteSharedData.ExpiresAt != nil && noteSharedData.ExpiresAt.Before(time.Now()) {
		if err := biz.noteShareLinkRepo.DeleteNoteShareLink(ctx, int64(noteSharedData.Id)); err != nil {
			return 0, "", core.ErrInternalServerError.
				WithError(noteShareLinkEntity.ErrCannotDeleteShareLink.Error()).
				WithDebug(err.Error())
		}

		return 0, "", core.ErrBadRequest.WithError("share link expired")
	}

	noteId := int(noteSharedData.NoteID)
	permission := noteSharedData.Permission

	value := fmt.Sprintf("%d|%s", noteId, permission)
	var ttl time.Duration = 24 * time.Hour
	if noteSharedData.ExpiresAt != nil {
		expirySeconds := int(time.Until(*noteSharedData.ExpiresAt).Seconds())
		if expirySeconds > 0 {
			ttl = min(ttl, time.Duration(expirySeconds)*time.Second)
		}
	}

	if setErr := biz.redisClient.Set(ctx, key, value, ttl); setErr != nil {
		return 0, "", core.ErrInternalServerError.
			WithError(noteShareLinkEntity.ErrCannotCreateShareLink.Error()).
			WithDebug(setErr.Error())
	}

	return noteId, permission, nil
}
