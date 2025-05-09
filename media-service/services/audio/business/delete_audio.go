package business

import (
	"context"
	"fmt"
	"strings"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteAudio(ctx context.Context, id int) error {
	// Get media data, without extra infos
	audio, err := biz.audioRepo.GetAudioById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetAudio.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAudio.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	note, err := biz.noteRepo.GetNoteById(ctx, int(audio.NoteID))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetNoteByID.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNoteByID.Error()).
			WithDebug(err.Error())
	}

	if note.UserId != int64(requesterId) {
		return core.ErrInternalServerError.
			WithError(entity.ErrRequesterIsNotOwner.Error())
	}

	if err := biz.audioRepo.DeleteAudio(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteAudio.Error()).
			WithDebug(err.Error())
	}

	urlParts := strings.Split(audio.FileURL, "/audios/")
	audioId := urlParts[len(urlParts)-1]
	fileKey := fmt.Sprintf("audios/%s", audioId)
	if err := biz.s3Client.DeleteObject(ctx, fileKey); err != nil {
		fmt.Printf("Failed to delete file from S3: %v\n", err)
	}

	return nil
}
