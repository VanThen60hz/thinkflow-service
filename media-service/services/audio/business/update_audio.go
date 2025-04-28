package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error {
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

	hasWritePermission, err := biz.collabRepo.HasReadPermission(ctx, int(audio.NoteID), requesterId)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetPermission.Error()).
			WithDebug(err.Error())
	}
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

	if note.UserId != int64(requesterId) && !hasWritePermission {
		return core.ErrInternalServerError.
			WithError(entity.ErrRequesterCannotModify.Error())
	}

	if err := biz.audioRepo.UpdateAudio(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateAudio.Error()).
			WithDebug(err.Error())
	}

	return nil
}
