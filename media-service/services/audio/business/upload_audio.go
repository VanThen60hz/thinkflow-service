package business

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"thinkflow-service/helper"
	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

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
		return nil, core.ErrInternalServerError.
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

	return &data, nil
}
