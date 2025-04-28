package business

import (
	"context"
	"fmt"
	"strings"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteAttachment(ctx context.Context, id int64) error {
	attachment, err := biz.attachmentRepo.GetByID(ctx, id)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if err := biz.attachmentRepo.DeleteAttachment(ctx, int(id)); err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	urlParts := strings.Split(attachment.FileURL, "/attachments/")
	attachmentId := urlParts[len(urlParts)-1]
	fileKey := fmt.Sprintf("attachments/%s", attachmentId)
	if err := biz.s3Client.DeleteObject(ctx, fileKey); err != nil {
		fmt.Printf("Failed to delete file from S3: %v\n", err)
	}

	return nil
}
