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

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) AcceptSharedNote(ctx context.Context, token string) (int, error) {
	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return 0, core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return 0, core.ErrInternalServerError.WithError("invalid requester type in context")
	}

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	noteId, permission, err := biz.getNoteShareData(ctx, token)
	if err != nil {
		return 0, err
	}

	if _, err := biz.noteRepo.GetNoteById(ctx, noteId); err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return 0, core.ErrNotFound.WithError("note not found")
		}
		return 0, core.ErrInternalServerError.WithError("cannot get note").WithDebug(err.Error())
	}

	newCollabData := &collabEntity.CollaborationCreation{
		UserId:     requesterId,
		NoteId:     noteId,
		Permission: collabEntity.PermissionType(permission),
	}

	if err := biz.collabRepo.AddNewCollaboration(ctx, newCollabData); err != nil {
		return 0, core.ErrInternalServerError.WithError("cannot add collaboration").WithDebug(err.Error())
	}

	return noteId, nil
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

	if !errors.Is(err, core.ErrRecordNotFound) {
		return 0, "", core.ErrInternalServerError.WithError("cannot get share link").WithDebug(err.Error())
	}

	noteSharedData, err := biz.noteShareLinkRepo.GetNoteShareLinkByToken(ctx, token)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return 0, "", core.ErrNotFound.WithError("share link not found")
		}
		return 0, "", core.ErrInternalServerError.WithError("cannot get share link").WithDebug(err.Error())
	}

	if noteSharedData.ExpiresAt != nil && noteSharedData.ExpiresAt.Before(time.Now()) {
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
