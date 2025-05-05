package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/gorm"
)

type SummaryNoteResponse struct {
	Summary string `json:"summary"`
}

type SummaryNoteFromTextRequest struct {
	Text              string   `json:"text"`
	ListTranscripText []string `json:"list_transcrip_text"`
}

type SummaryNoteFromTextResponse struct {
	SummaryFromAllText string `json:"summary_from_all_text"`
}

type SummaryNoteFromAudioResponse struct {
	SummaryAudiosAndText string `json:"summary_audios_and_text"`
}

func (biz *business) SummaryNote(ctx context.Context, noteID int) (*SummaryNoteResponse, error) {
	_, err := biz.noteRepo.GetNoteById(ctx, noteID)
	if err != nil {
		return nil, err
	}

	text := getTextOrEmpty(ctx, biz, noteID)
	audios := getAudiosOrEmpty(ctx, biz, noteID)

	summaryText, err := biz.generateSummary(ctx, text, audios)
	if err != nil {
		return nil, err
	}

	summaryResp := &SummaryNoteResponse{Summary: summaryText}

	summaryID, err := biz.summaryRepo.CreateSummary(ctx, summaryText)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to create summary")
	}

	updateData := &entity.NoteDataUpdate{SummaryID: &summaryID}
	if err := biz.noteRepo.UpdateNote(ctx, noteID, updateData); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to update note")
	}

	return summaryResp, nil
}

func (biz *business) generateSummary(ctx context.Context, text string, audios []*pb.PublicAudioInfo) (string, error) {
	if len(audios) == 0 {
		return biz.callTextSummaryAPI(ctx, text, []string{})
	}

	allTranscribed := true
	var transcriptTexts []string
	for _, audio := range audios {
		if audio.TranscriptId == 0 {
			allTranscribed = false
			break
		}
		transcript, err := biz.transcriptRepo.GetTranscriptById(ctx, audio.TranscriptId)
		if err != nil {
			return "", core.ErrInternalServerError.WithError("failed to get transcript")
		}
		transcriptTexts = append(transcriptTexts, transcript.Content)
	}

	if allTranscribed {
		return biz.callTextSummaryAPI(ctx, text, transcriptTexts)
	}
	return biz.callAudioSummaryAPI(ctx, text, audios)
}

func (biz *business) callTextSummaryAPI(ctx context.Context, text string, transcripts []string) (string, error) {
	reqBody := SummaryNoteFromTextRequest{Text: text, ListTranscripText: transcripts}
	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post("http://ai-service:5000/summary_note_from_all_text", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", core.ErrInternalServerError.WithError("failed to call summary API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", core.ErrInternalServerError.WithError(fmt.Sprintf("summary API error: %d - %s", resp.StatusCode, string(bodyBytes)))
	}

	var result SummaryNoteFromTextResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", core.ErrInternalServerError.WithError("failed to parse summary response")
	}
	return result.SummaryFromAllText, nil
}

func (biz *business) callAudioSummaryAPI(ctx context.Context, text string, audios []*pb.PublicAudioInfo) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("text", text)

	for _, audio := range audios {
		key, err := extractS3KeyFromURL(audio.FileUrl)
		if err != nil {
			return "", core.ErrInternalServerError.WithError("invalid file URL")
		}
		audioFile, err := biz.s3Client.Download(ctx, key)
		if err != nil {
			return "", core.ErrInternalServerError.WithError(err.Error())
		}
		defer audioFile.Close()

		part, err := writer.CreateFormFile("files", key)
		if err != nil {
			return "", core.ErrInternalServerError.WithError("failed to create form file")
		}
		if _, err := io.Copy(part, audioFile); err != nil {
			return "", core.ErrInternalServerError.WithError("failed to copy file content")
		}
	}
	writer.Close()

	resp, err := http.Post("http://ai-service:5000/summary_note_from_all_audio_and_text", writer.FormDataContentType(), body)
	if err != nil {
		return "", core.ErrInternalServerError.WithError("failed to call summary API")
	}
	defer resp.Body.Close()

	var result SummaryNoteFromAudioResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", core.ErrInternalServerError.WithError("failed to parse summary response")
	}
	return result.SummaryAudiosAndText, nil
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

func getAudiosOrEmpty(ctx context.Context, biz *business, noteID int) []*pb.PublicAudioInfo {
	audios, err := biz.audioRepo.GetAudiosByNoteId(ctx, int64(noteID))
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("Error fetching audios:", err)
	}
	return audios
}

func extractS3KeyFromURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(parsed.Path, "/"), nil
}
