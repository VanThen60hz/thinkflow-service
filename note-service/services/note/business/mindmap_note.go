package business

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"thinkflow-service/proto/pb"
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

	var mindmapResp datatypes.JSON

	if note.SummaryID != nil {
		summary, err := biz.summaryRepo.GetSummaryById(ctx, *note.SummaryID)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to get summary")
		}

		mindmapResp, err = callMindmapAPI("http://ai-service:5000/mindmap_note_from_summary", map[string]string{
			"summary": summary.SummaryText,
		})
		if err != nil {
			return nil, err
		}
	} else {
		mindmapResp, err = biz.generateMindmapFromTextAndAudio(ctx, noteID)
		if err != nil {
			return nil, err
		}
	}

	return biz.saveMindmap(ctx, noteID, mindmapResp)
}

func (biz *business) generateMindmapFromTextAndAudio(ctx context.Context, noteID int) (datatypes.JSON, error) {
	text := getTextOrEmpty(ctx, biz, noteID)

	audios := getAudiosOrEmpty(ctx, biz, noteID)

	if text == "" && len(audios) == 0 {
		return nil, core.ErrBadRequest.WithError("cannot generate mindmap: both text and audio are empty")
	}

	allTranscribed := true
	transcriptTexts := []string{}
	for _, audio := range audios {
		if audio.TranscriptId == 0 {
			allTranscribed = false
			break
		}
		t, err := biz.transcriptRepo.GetTranscriptById(ctx, audio.TranscriptId)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to get transcript")
		}
		transcriptTexts = append(transcriptTexts, t.Content)
	}

	var err error
	var mindmapResp datatypes.JSON
	if allTranscribed {
		if len(transcriptTexts) == 0 {
			transcriptTexts = []string{""}
		}

		mindmapResp, err = callMindmapAPI("http://ai-service:5000/mindmap_note_from_all_text", MindmapNoteFromTextRequest{
			Text:              text,
			ListTranscripText: transcriptTexts,
		})
		if err != nil {
			return nil, err
		}
	} else {
		mindmapResp, err = biz.callMindmapWithAudioAndText(ctx, text, audios)
		if err != nil {
			return nil, err
		}
	}

	return mindmapResp, nil
}

func (biz *business) saveMindmap(ctx context.Context, noteID int, mindmapResp datatypes.JSON) (datatypes.JSON, error) {
	mindmapData, _ := json.Marshal(mindmapResp)
	mindmapID, err := biz.mindmapRepo.CreateMindmap(ctx, string(mindmapData))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to save mindmap")
	}

	if err := biz.noteRepo.UpdateNote(ctx, noteID, &entity.NoteDataUpdate{MindmapID: &mindmapID}); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to update note with mindmap ID")
	}

	return mindmapResp, nil
}

func callMindmapAPI(url string, payload interface{}) (datatypes.JSON, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to marshal request")
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to call API: " + url)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to read API response")
	}

	return datatypes.JSON(data), nil
}

func (biz *business) callMindmapWithAudioAndText(ctx context.Context, text string, audios []*pb.PublicAudioInfo) (datatypes.JSON, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("text", text); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to write text field")
	}

	for _, audio := range audios {
		key, err := extractS3KeyFromURL(audio.FileUrl)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("invalid file URL")
		}

		file, err := biz.s3Client.Download(ctx, key)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError(err.Error())
		}
		defer file.Close()

		part, err := writer.CreateFormFile("files", key)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to create form file")
		}

		if _, err := io.Copy(part, file); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to copy file content")
		}
	}

	writer.Close()

	resp, err := http.Post("http://ai-service:5000/mindmap_note_from_all_audio_and_text", writer.FormDataContentType(), body)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to call audio+text mindmap API")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to read response body")
	}

	return datatypes.JSON(data), nil
}
