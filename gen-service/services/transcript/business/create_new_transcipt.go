package business

import (
	"context"

	"thinkflow-service/services/transcript/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewTranscript(ctx context.Context, data *entity.TranscriptDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // transcript user id, id of who creates this new Transcript

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.transcriptRepo.AddNewTranscript(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateTranscript.Error())
	}

	return nil
}
