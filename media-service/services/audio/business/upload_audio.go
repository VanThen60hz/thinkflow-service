package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"thinkflow-service/helper"
	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) transcribeAudio(ctx context.Context, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		return "", err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return "", err
	}
	if _, err := part.Write(buf.Bytes()); err != nil {
		return "", err
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", "http://ai-service:5000/transcript_from_an_audio", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type transcriptResponse struct {
		Transcript string `json:"transcript"`
	}

	var result transcriptResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}

	return result.Transcript, nil
}

func (biz *business) UploadAudio(ctx context.Context, tempFile string, file *multipart.FileHeader, noteID int64) (*entity.AudioDataCreation, error) {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasWritePermission, err := biz.collabRepo.HasReadPermission(ctx, int(noteID), requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetPermission.Error()).
			WithDebug(err.Error())
	}
	note, err := biz.noteRepo.GetNoteById(ctx, int(noteID))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithError(entity.ErrCannotGetNoteByID.Error()).
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNoteByID.Error()).
			WithDebug(err.Error())
	}

	if note.UserId != int64(requesterId) && !hasWritePermission {
		return nil, core.ErrBadRequest.
			WithError(entity.ErrRequesterCannotModify.Error())
	}

	processor := helper.NewMediaProcessor()
	audioInfo, err := processor.ProcessAudio(file)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateAudio.Error()).
			WithDebug(err.Error())
	}

	fileUrl, err := biz.s3Client.Upload(ctx, tempFile, "audios")
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateAudio.Error()).
			WithDebug(err.Error())
	}

	data := entity.AudioDataCreation{
		NoteID:  noteID,
		FileURL: fileUrl,
		Format:  audioInfo.Format,
	}

	if err := biz.audioRepo.AddNewAudio(ctx, &data); err != nil {
		urlParts := strings.Split(fileUrl, "/audios/")
		audioId := urlParts[len(urlParts)-1]
		fileKey := fmt.Sprintf("audios/%s", audioId)
		if err := biz.s3Client.DeleteObject(ctx, fileKey); err != nil {
			fmt.Printf("Failed to delete file from S3: %v\n", err)
		}
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateAudio.Error()).
			WithDebug(err.Error())
	}

	go func() {
		transcriptCtx := context.Background()

		transcript, err := biz.transcribeAudio(transcriptCtx, file)
		if err != nil {
			fmt.Printf("Failed to transcribe audio: %v\n", err)
			return
		}

		transcriptID, err := biz.transcriptRepo.CreateTranscript(transcriptCtx, transcript)
		if err != nil {
			fmt.Printf("Failed to create transcript: %v\n", err)
			return
		}

		if err := biz.audioRepo.UpdateAudio(transcriptCtx, data.Id, &entity.AudioDataUpdate{
			TranscriptID: &transcriptID,
		}); err != nil {
			fmt.Printf("Failed to update audio with transcript ID: %v\n", err)
			return
		}

		note, err := biz.noteRepo.GetNoteById(transcriptCtx, int(noteID))
		if err != nil {
			fmt.Printf("Failed to get note info for notification: %v\n", err)
			return
		}

		biz.sendNotificationToAudioMembers(transcriptCtx, note, requesterId, "TRANSCRIPT_GENERATED", fmt.Sprintf("Audio in note '%s' has been transcribed", note.Title))
	}()

	return &data, nil
}
