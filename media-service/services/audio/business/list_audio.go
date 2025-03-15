package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListAudios(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Audio, error) {
	audios, err := biz.audioRepo.ListAudios(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListAudio.Error()).
			WithDebug(err.Error())
	}

	// Get extra infos: User
	userIds := make([]int, len(audios))

	for i := range userIds {
		userIds[i] = audios[i].UserId
	}

	return audios, nil
}
