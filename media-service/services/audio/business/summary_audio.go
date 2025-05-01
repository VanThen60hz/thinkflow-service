package business

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

type SummaryResponse struct {
	Summary string `json:"summary"`
}

type AudioSummaryResponse struct {
	SummaryFromAudios string  `json:"summary_from_audios"`
	TimeTakenSec      float64 `json:"time_taken_sec"`
}

func (biz *business) SummaryAudio(ctx context.Context, audioID int) (*SummaryResponse, error) {
	audio, err := biz.audioRepo.GetAudioById(ctx, audioID)
	if err != nil {
		return nil, err
	}

	var summaryResp *SummaryResponse
	var summaryText string

	if audio.TranscriptID != nil {
		transcript, err := biz.transcriptRepo.GetTranscriptById(ctx, *audio.TranscriptID)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to get transcript")
		}

		// Call summary API with transcript text
		reqBody := map[string]string{
			"text": transcript.Content,
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

		if err := json.NewDecoder(resp.Body).Decode(&summaryResp); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to parse summary response")
		}
		summaryText = summaryResp.Summary
	} else {
		// Download audio file from S3
		key, err := extractS3KeyFromURL(audio.FileURL)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("invalid file URL")
		}

		audioFile, err := biz.s3Client.Download(ctx, key)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError(err.Error())
		}
		defer audioFile.Close()

		// Create a temporary file to store the audio
		tempFile, err := os.CreateTemp("", "audio-*.mp3")
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to create temporary file")
		}
		defer os.Remove(tempFile.Name())
		defer tempFile.Close()

		// Copy the audio content to the temporary file
		if _, err := io.Copy(tempFile, audioFile); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to copy audio content")
		}

		// Create multipart form data
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("files", tempFile.Name())
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to create form file")
		}

		// Reset the file pointer to the beginning
		if _, err := tempFile.Seek(0, 0); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to seek file")
		}

		_, err = io.Copy(part, tempFile)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to copy file content")
		}
		writer.Close()

		resp, err := http.Post(
			"http://ai-service:5000/summary_from_all_audios",
			writer.FormDataContentType(),
			body,
		)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to call summary API")
		}
		defer resp.Body.Close()

		var audioSummaryResp AudioSummaryResponse
		if err := json.NewDecoder(resp.Body).Decode(&audioSummaryResp); err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to parse summary response")
		}
		summaryText = audioSummaryResp.SummaryFromAudios
		summaryResp = &SummaryResponse{Summary: summaryText}
	}

	summaryId, err := biz.summaryRepo.CreateSummary(ctx, summaryText)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to create summary")
	}

	updateData := &entity.AudioDataUpdate{
		FileURL:      &audio.FileURL,
		TranscriptID: audio.TranscriptID,
		SummaryID:    &summaryId,
	}

	if err := biz.audioRepo.UpdateAudio(ctx, audioID, updateData); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed to update audio")
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
