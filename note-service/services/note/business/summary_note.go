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

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
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
	note, err := biz.noteRepo.GetNoteById(ctx, noteID)
	if err != nil {
		return nil, err
	}
	fmt.Println(note)

	text, err := biz.textRepo.GetTextByNoteId(ctx, noteID)
	if err != nil {
		return nil, err
	}
	fmt.Println(text)

	audios, err := biz.audioRepo.GetAudiosByNoteId(ctx, int64(noteID))
	fmt.Println(audios)
	if err != nil {
		return nil, err
	}

	var summaryResp *SummaryNoteResponse
	var summaryText string

	allTranscribed := true
	var transcriptTexts []string

	for _, audio := range audios {
		if audio.TranscriptId == 0 {
			allTranscribed = false
			break
		}

		transcript, err := biz.transcriptRepo.GetTranscriptById(ctx, audio.TranscriptId)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to get transcript")
		}
		transcriptTexts = append(transcriptTexts, transcript.Content)
	}

	if allTranscribed {
		// Case 1: All audios are transcribed
		reqBody := SummaryNoteFromTextRequest{
			Text:              text.TextString,
			ListTranscripText: transcriptTexts,
		}
		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to marshal request body")
		}

		resp, err := http.Post(
			"http://ai-service:5000/mindmap_note_from_all_text",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to call mindmap API")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return nil, core.ErrInternalServerError.WithError(
				fmt.Sprintf("mindmap API error: %d - %s", resp.StatusCode, string(bodyBytes)),
			)
		}

		var textSummaryResp SummaryNoteFromTextResponse
		if err := json.NewDecoder(resp.Body).Decode(&textSummaryResp); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to parse summary response")
		}
		summaryText = textSummaryResp.SummaryFromAllText
	} else {
		// Case 2: Not all audios are transcribed
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Add text field
		if err := writer.WriteField("text", text.TextString); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to write text field")
		}

		// Add audio files
		for _, audio := range audios {
			key, err := extractS3KeyFromURL(audio.FileUrl)
			if err != nil {
				return nil, core.ErrInternalServerError.WithError("invalid file URL")
			}

			audioFile, err := biz.s3Client.Download(ctx, key)
			if err != nil {
				return nil, core.ErrInternalServerError.WithError(err.Error())
			}
			defer audioFile.Close()

			part, err := writer.CreateFormFile("files", key)
			if err != nil {
				return nil, core.ErrInternalServerError.WithError("failed to create form file")
			}

			if _, err := io.Copy(part, audioFile); err != nil {
				return nil, core.ErrInternalServerError.WithError("failed to copy file content")
			}
		}
		writer.Close()

		resp, err := http.Post(
			"http://ai-service:5000/summary_note_from_all_audio_and_text",
			writer.FormDataContentType(),
			body,
		)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to call summary API")
		}
		defer resp.Body.Close()

		var audioSummaryResp SummaryNoteFromAudioResponse
		if err := json.NewDecoder(resp.Body).Decode(&audioSummaryResp); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to parse summary response")
		}
		summaryText = audioSummaryResp.SummaryAudiosAndText
	}

	summaryResp = &SummaryNoteResponse{Summary: summaryText}

	summaryId, err := biz.summaryRepo.CreateSummary(ctx, summaryText)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to create summary")
	}

	updateData := &entity.NoteDataUpdate{
		SummaryID: &summaryId,
	}

	if err := biz.noteRepo.UpdateNote(ctx, noteID, updateData); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to update note")
	}

	return summaryResp, nil
}

func extractS3KeyFromURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(parsed.Path, "/"), nil
}
