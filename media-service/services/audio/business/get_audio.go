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

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasReadPermission, err := biz.collabRepo.HasReadPermission(ctx, int(data.NoteID), requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetPermission.Error()).
			WithDebug(err.Error())
	}
	note, err := biz.noteRepo.GetNoteById(ctx, int(data.NoteID))
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

	if data.TranscriptID != nil {
		transcript, err := biz.transcriptRepo.GetTranscriptById(ctx, *data.TranscriptID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetTranscript.Error()).
				WithDebug(err.Error())
		}
		data.Transcript = transcript
	}

	if data.SummaryID != nil {
		summary, err := biz.summaryRepo.GetSummaryById(ctx, *data.SummaryID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetSummary.Error()).
				WithDebug(err.Error())
		}
		data.Summary = summary
	}

	return data, nil
}
