package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetAudiosByNoteIdInt64(ctx context.Context, noteId int64) ([]entity.Audio, error) {
	filter := &entity.Filter{
		NoteID: &noteId,
	}

	audios, err := biz.audioRepo.ListAudios(ctx, filter, nil)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListAudio.Error()).
			WithDebug(err.Error())
	}

	return audios, nil
}
