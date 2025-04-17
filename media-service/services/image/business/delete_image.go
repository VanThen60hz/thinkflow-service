package business

import (
	"context"
	"fmt"
	"strings"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteImage(ctx context.Context, id int) error {
	image, err := biz.imageRepo.GetImageById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetImage.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetImage.Error()).
			WithDebug(err.Error())
	}

	if err := biz.imageRepo.DeleteImage(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteImage.Error()).
			WithDebug(err.Error())
	}

	urlParts := strings.Split(image.Url, "/images/")
	imageId := urlParts[len(urlParts)-1]
	fileKey := fmt.Sprintf("images/%s", imageId)

	if err := biz.s3Client.DeleteObject(ctx, fileKey); err != nil {
		fmt.Printf("Failed to delete file from S3: %v\n", err)
	}

	return nil
}
