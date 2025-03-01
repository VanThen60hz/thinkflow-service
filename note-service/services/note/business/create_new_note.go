package business

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewNote(ctx context.Context, data *entity.NoteDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // note user id, id of who creates this new note

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.noteRepo.AddNewNote(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateNote.Error())
	}

	return nil
}
