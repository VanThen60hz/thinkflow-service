package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetAudioById(ctx context.Context, id int) (*entity.Audio, error) {
	data, err := biz.audioRepo.GetAudioById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAudio.Error()).
			WithDebug(err.Error())
	}

	transcript, err := biz.transcriptRepo.GetTranscriptById(ctx, *data.TranscriptID)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetTranscript.Error()).
			WithDebug(err.Error())
	}
	data.Transcript = transcript

	summary, err := biz.summaryRepo.GetSummaryById(ctx, *data.SummaryID)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetSummary.Error()).
			WithDebug(err.Error())
	}
	data.Summary = summary

	return data, nil
}
