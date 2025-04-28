package business

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"thinkflow-service/common"
	noteEntity "thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/core"
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

	link, err := biz.CreateNoteShareLink(ctx, noteId, permission, expiresAt)
	if err != nil {
		return err
	}

	if link == nil {
		return core.ErrInternalServerError.WithError("cannot create share link")
	}

	clientURL := os.Getenv("CLIENT_URL")
	clientURL = strings.TrimRight(clientURL, "/")
	shareURL := fmt.Sprintf("%s/share/%s", clientURL, link.Token)

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

	if err := biz.emailService.SendGenericLink(ctx, email, "ThinkFlow Note Share Invitation", emailData); err != nil {
		return core.ErrInternalServerError.WithError("cannot send email").WithDebug(err.Error())
	}

	return nil
}
