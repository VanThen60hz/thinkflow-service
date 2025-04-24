package business

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListNotesSharedWithMe(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error) {
	uID, err := core.FromBase58(*filter.UserId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListNote.Error()).
			WithDebug(err.Error())
	}

	userId := int(uID.GetLocalID())

	// Get collaborations with paging
	collaborations, err := biz.collabRepo.GetCollaborationByUserId(ctx, userId, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListNote.Error()).
			WithDebug(err.Error())
	}

	var sharedNotes []entity.Note
	for i := range collaborations {
		if collaborations[i].NoteId == 0 {
			continue
		}

		collabNote, err := biz.noteRepo.GetNoteById(ctx, collaborations[i].NoteId)
		if err != nil {
			if err == core.ErrRecordNotFound {
				continue
			}
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetNote.Error()).
				WithDebug(err.Error())
		}

		if collabNote.Archived {
			continue
		}

		// Set permission based on collaboration
		if collaborations[i].Permission == "write" {
			collabNote.Permission = "write"
		} else {
			collabNote.Permission = "read"
		}

		sharedNotes = append(sharedNotes, *collabNote)
	}

	// Get extra infos: User
	userIds := make([]int, len(sharedNotes))
	for i := range sharedNotes {
		userIds[i] = sharedNotes[i].UserId
	}

	users, err := biz.userRepo.GetUsersByIds(ctx, userIds)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListNote.Error()).
			WithDebug(err.Error())
	}

	mUser := make(map[int]core.SimpleUser, len(users))
	for i := range users {
		mUser[users[i].Id] = users[i]
	}

	for i := range sharedNotes {
		if user, ok := mUser[sharedNotes[i].UserId]; ok {
			sharedNotes[i].User = &user
		}
	}

	if len(sharedNotes) > 0 {
		sharedNotes[len(sharedNotes)-1].Mask()
		paging.NextCursor = sharedNotes[len(sharedNotes)-1].FakeId.String()
	}

	return sharedNotes, nil
}
