package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/datatypes"
)

type MindmapNoteFromTextRequest struct {
	Text              string   `json:"text"`
	ListTranscripText []string `json:"list_transcrip_text"`
}

func (biz *business) MindmapNote(ctx context.Context, noteID int) (datatypes.JSON, error) {
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

	var mindmapResp datatypes.JSON

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
		reqBody := MindmapNoteFromTextRequest{
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

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to read mindmap response body")
		}
		mindmapResp = datatypes.JSON(bodyBytes)
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
			"http://ai-service:5000/mindmap_note_from_all_audio_and_text",
			writer.FormDataContentType(),
			body,
		)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to call mindmap API")
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to read mindmap response body")
		}
		mindmapResp = datatypes.JSON(bodyBytes)
	}

	mindmapData, err := json.Marshal(mindmapResp)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to marshal mindmap data")
	}

	mindmapId, err := biz.mindmapRepo.CreateMindmap(ctx, string(mindmapData))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	updateData := &entity.NoteDataUpdate{
		MindmapID: &mindmapId,
	}

	if err := biz.noteRepo.UpdateNote(ctx, noteID, updateData); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to update note")
	}

	return mindmapResp, nil
}
