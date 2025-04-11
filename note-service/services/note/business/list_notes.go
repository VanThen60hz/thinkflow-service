package business

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListNotes(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error) {
	notes, err := biz.noteRepo.ListNotes(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListNote.Error()).
			WithDebug(err.Error())
	}

	uID, err := core.FromBase58(*filter.UserId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListNote.Error()).
			WithDebug(err.Error())
	}

	collaborations, err := biz.collabRepo.GetCollaborationByUserId(ctx, int(uID.GetLocalID()))
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListNote.Error()).
			WithDebug(err.Error())
	}

	for i := range collaborations {
		if collaborations[i].NoteId == 0 {
			continue
		}
		collabNote, err := biz.noteRepo.GetNoteById(ctx, collaborations[i].NoteId)
		if err != nil {
			if err == core.ErrRecordNotFound {
				return nil, core.ErrNotFound.
					WithDebug(err.Error())
			}
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetNote.Error()).
				WithDebug(err.Error())
		}

		if collabNote.Archived {
			continue
		}

		notes = append(notes, *collabNote)
	}

	// Get extra infos: User
	userIds := make([]int, len(notes))

	for i := range userIds {
		userIds[i] = notes[i].UserId
	}

	users, err := biz.userRepo.GetUsersByIds(ctx, userIds)

	mUser := make(map[int]core.SimpleUser, len(users))
	for i := range users {
		mUser[users[i].Id] = users[i]
	}

	for i := range notes {
		if user, ok := mUser[notes[i].UserId]; ok {
			notes[i].User = &user
		}
	}

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListNote.Error()).
			WithDebug(err.Error())
	}

	return notes, nil
}
