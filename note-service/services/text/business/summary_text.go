package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/gorm"
)

type SummaryResponse struct {
	Summary string `json:"summary"`
}

func (biz *business) SummaryText(ctx context.Context, textID int) (*SummaryResponse, error) {
	text, err := biz.GetTextById(ctx, textID)
	if err != nil {
		return nil, err
	}
	if text == nil {
		return nil, core.ErrNotFound.WithError("text not found")
	}

	// Get note info to send notification
	note, err := biz.noteRepo.GetNoteById(ctx, int(text.NoteID))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to get note info")
	}
	if note == nil {
		return nil, core.ErrNotFound.WithError("note not found")
	}

	textString := getTextOrEmpty(ctx, biz, textID)

	reqBody := map[string]string{
		"text": textString,
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
		TextString:  textString,
		SummaryID:   &summaryId,
	}

	if err := biz.UpdateText(ctx, text.Id, updateData); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to update text")
	}

	// Get requester info
	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return nil, core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return nil, core.ErrInternalServerError.WithError("invalid requester type in context")
	}

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	// Send notifications
	biz.sendNotificationToNoteMembers(ctx, note, requesterId, "TEXT_SUMMARY_GENERATED", fmt.Sprintf("Text in note '%s' has been summarized", note.Title))

	return &summaryResp, nil
}

func getTextOrEmpty(ctx context.Context, biz *business, noteID int) string {
	text, err := biz.textRepo.GetTextByNoteId(ctx, noteID)
	if err == nil && text != nil {
		return text.TextString
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("Error fetching text:", err)
	}
	return ""
}
