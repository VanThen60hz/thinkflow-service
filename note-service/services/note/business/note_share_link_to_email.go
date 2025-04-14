package business

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"thinkflow-service/common"
	noteShareLinkEntity "thinkflow-service/services/note-share-links/entity"
	noteEntity "thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/core"
	"github.com/google/uuid"
)

func (biz *business) NoteShareLinkToEmail(ctx context.Context, noteId int64, email, permission string, expiresAt *time.Time) error {
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

	noteData, err := biz.noteRepo.GetNoteById(ctx, int(noteId))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.WithError(noteEntity.ErrNoteNotFound.Error())
		}
		return core.ErrInternalServerError.WithError(noteEntity.ErrCannotGetNote.Error()).WithDebug(err.Error())
	}

	if requesterId != noteData.UserId {
		return core.ErrForbidden.WithError(noteEntity.ErrRequesterCannotModify.Error())
	}

	tid := uuid.New().String()
	uid = core.NewUID(uint32(noteId), 1, 1)
	sub := uid.String()

	token, _, err := biz.jwtProvider.IssueToken(ctx, tid, sub)
	if err != nil {
		return core.ErrInternalServerError.WithError(noteShareLinkEntity.ErrCannotCreateShareLink.Error()).WithDebug(err.Error())
	}

	data := &noteShareLinkEntity.NoteShareLinkCreation{
		NoteID:     noteId,
		Permission: permission,
		Token:      token,
		ExpiresAt:  expiresAt,
	}

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.noteShareLinkRepo.AddNewNoteShareLink(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(noteShareLinkEntity.ErrCannotCreateShareLink.Error()).WithDebug(err.Error())
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
		return core.ErrInternalServerError.WithError(noteShareLinkEntity.ErrCannotCreateShareLink.Error()).WithDebug(err.Error())
	}

	clientURL := os.Getenv("CLIENT_URL")
	clientURL = strings.TrimRight(clientURL, "/")
	shareURL := fmt.Sprintf("%s/share/%s", clientURL, token)

	var expireMinutes *int
	if expiresAt != nil {
		minutes := int(time.Until(*expiresAt).Minutes())
		if minutes > 0 {
			expireMinutes = &minutes
		}
	}

	buttonText := "View Note"
	if permission == "write" {
		buttonText = "Edit Note"
	}

	emailData := emailc.LinkMailData{
		Title:         "Note Share Link",
		UserEmail:     email,
		MessageIntro:  "You have been invited to access a note on ThinkFlow. Use the link below to view or edit the note.",
		Link:          shareURL,
		ButtonText:    buttonText,
		LinkTypeDesc:  "note share link",
		ExpireMinutes: expireMinutes,
	}

	fmt.Println("emailData:", emailData)

	if err := biz.emailService.SendGenericLink(ctx, email, "ThinkFlow Note Share Invitation", emailData); err != nil {
		return core.ErrInternalServerError.WithError(noteEntity.ErrCannotSendShareLinkEmail.Error()).WithDebug(err.Error())
	}

	return nil
}
