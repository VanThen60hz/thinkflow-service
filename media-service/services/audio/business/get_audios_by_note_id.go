package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetAudiosByNoteId(ctx context.Context, noteId int) ([]entity.Audio, error) {
	noteId64 := int64(noteId)
	filter := &entity.Filter{
		NoteID: &noteId64,
	}
	paging := &core.Paging{
		Page:  1,
		Limit: 100,
	}

	audios, err := biz.audioRepo.ListAudios(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListAudio.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasReadPermission, err := biz.collabRepo.HasReadPermission(ctx, int(noteId), requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetPermission.Error()).
			WithDebug(err.Error())
	}
	note, err := biz.noteRepo.GetNoteById(ctx, int(noteId))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithError(entity.ErrCannotGetNoteByID.Error()).
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNoteByID.Error()).
			WithDebug(err.Error())
	}

	if note.UserId != int64(requesterId) && !hasReadPermission {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrRequesterCannotRead.Error())
	}

	return audios, nil
}
