package api

import (
	"time"

	"thinkflow-service/services/note/entity"
)

type NoteResponse struct {
	*entity.Note
	FormattedCreatedAt string `json:"created_at"`
	FormattedUpdatedAt string `json:"updated_at"`
}

func NewNoteResponse(note *entity.Note) *NoteResponse {
	if note == nil {
		return nil
	}

	var createdAt, updatedAt time.Time
	if note.CreatedAt != nil {
		createdAt = *note.CreatedAt
	} else {
		createdAt = time.Time{}
	}
	if note.UpdatedAt != nil {
		updatedAt = *note.UpdatedAt
	} else {
		updatedAt = time.Time{}
	}

	return &NoteResponse{
		Note:               note,
		FormattedCreatedAt: createdAt.Format("02/01/2006"),
		FormattedUpdatedAt: updatedAt.Format("02/01/2006"),
	}
}

type ListNoteResponse struct {
	Data   []NoteResponse `json:"data"`
	Paging interface{}    `json:"paging,omitempty"`
}
