package business

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"thinkflow-service/common"
	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UploadAttachment(ctx context.Context, tempFile string, file *multipart.FileHeader, noteID int64) (*entity.AttachmentCreation, error) {
	fileUrl, err := biz.s3Client.Upload(ctx, tempFile, "attachments")
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateAttachment.Error()).
			WithDebug(err.Error())
	}

	ext := filepath.Ext(file.Filename)

	attachment := &entity.AttachmentCreation{
		NoteID:    noteID,
		FileURL:   fileUrl,
		FileName:  file.Filename,
		Extension: ext,
		SizeBytes: file.Size,
		CloudName: common.KeyCompS3,
	}

	if err := biz.attachmentRepo.AddNewAttachment(ctx, attachment); err != nil {
		urlParts := strings.Split(fileUrl, "/attachments/")
		attachmentId := urlParts[len(urlParts)-1]
		fileKey := fmt.Sprintf("attachments/%s", attachmentId)
		if err := biz.s3Client.DeleteObject(ctx, fileKey); err != nil {
			fmt.Printf("Failed to delete file from S3: %v\n", err)
		}
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateAttachment.Error()).
			WithDebug(err.Error())
	}

	return attachment, nil
}
