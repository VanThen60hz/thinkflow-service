package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/text/entity"
)

type TextRepository interface {
	AddNewText(ctx context.Context, data *entity.TextDataCreation) error
	UpdateText(ctx context.Context, id int, data *entity.TextDataUpdate) error
	DeleteText(ctx context.Context, id int) error
	GetTextById(ctx context.Context, id int) (*entity.Text, error)
	GetTextByNoteId(ctx context.Context, noteId int) (*entity.Text, error)
}

type SummaryRepository interface {
	GetSummaryById(ctx context.Context, id int64) (*common.SimpleSummary, error)
}

type MindmapRepository interface {
	GetMindmapById(ctx context.Context, id int64) (*common.SimpleMindmap, error)
}

type business struct {
	textRepo    TextRepository
	summaryRepo SummaryRepository
	mindmapRepo MindmapRepository
}

func NewBusiness(
	textRepo TextRepository,
	summaryRepo SummaryRepository,
	mindmapRepo MindmapRepository,
) *business {
	return &business{
		textRepo:    textRepo,
		summaryRepo: summaryRepo,
		mindmapRepo: mindmapRepo,
	}
}
