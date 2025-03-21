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

	return audios, nil
}
