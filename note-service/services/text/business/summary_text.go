package business

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
)

type SummaryResponse struct {
	Summary string `json:"summary"`
}

func (biz *business) SummaryText(ctx context.Context, textID int) (*SummaryResponse, error) {
	text, err := biz.GetTextById(ctx, textID)
	if err != nil {
		return nil, err
	}

	reqBody := map[string]string{
		"text": text.TextString,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to marshal request body")
	}

	resp, err := http.Post(
		"http://ai-service:5000/summary_from_text",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to call summary API")
	}
	defer resp.Body.Close()

	var summaryResp SummaryResponse
	if err := json.NewDecoder(resp.Body).Decode(&summaryResp); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to parse summary response")
	}

	summaryId, err := biz.summaryRepo.CreateSummary(ctx, summaryResp.Summary)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to create summary")
	}

	updateData := &entity.TextDataUpdate{
		TextContent: text.TextContent,
		TextString:  text.TextString,
		SummaryID:   &summaryId,
	}

	if err := biz.UpdateText(ctx, text.Id, updateData); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to update text")
	}

	return &summaryResp, nil
}
