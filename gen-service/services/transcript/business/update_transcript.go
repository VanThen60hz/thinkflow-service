package business

import (
	"context"

	"thinkflow-service/services/transcript/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateTranscript(ctx context.Context, id int, data *entity.TranscriptDataUpdate) error {
	// Get Transcript data, without extra infos
	// transcript, err := biz.transcriptRepo.GetTranscriptById(ctx, id)
	_, err := biz.transcriptRepo.GetTranscriptById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetTranscript.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetTranscript.Error()).
			WithDebug(err.Error())
	}

	// requester := core.GetRequester(ctx)

	// uid, _ := core.FromBase58(requester.GetSubject())
	// requesterId := int(uid.GetLocalID())

	// // Only Transcript user can do this
	// if requesterId != transcript.UserId {
	// 	return core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	// }

	if err := biz.transcriptRepo.UpdateTranscript(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateTranscript.Error()).
			WithDebug(err.Error())
	}

	return nil
}
