package business

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetTextByNoteId(ctx context.Context, noteId int) (*entity.Text, error) {
	text, err := biz.textRepo.GetTextByNoteId(ctx, noteId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListText.Error()).
			WithDebug(err.Error())
	}

	if text.SummaryID != nil {
		summary, err := biz.summaryRepo.GetSummaryById(ctx, *text.SummaryID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetSummary.Error()).
				WithDebug(err.Error())
		}
		text.Summary = summary
	}

	if text.MindmapID != nil {
		mindmap, err := biz.mindmapRepo.GetMindmapById(ctx, *text.MindmapID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetMindmap.Error()).
				WithDebug(err.Error())
		}
		text.Mindmap = mindmap
	}

	return text, nil
}
