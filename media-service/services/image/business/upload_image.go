package business

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"thinkflow-service/common"
	"thinkflow-service/helper"
	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UploadImage(ctx context.Context, tempFile string, file *multipart.FileHeader) (*entity.ImageDataCreation, error) {
	processor := helper.NewMediaProcessor()
	imageInfo, err := processor.ProcessImage(file)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateImage.Error()).
			WithDebug(err.Error())
	}

	fileUrl, err := biz.s3Client.Upload(ctx, tempFile, "images")
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateImage.Error()).
			WithDebug(err.Error())
	}

	data := entity.ImageDataCreation{
		Url:       fileUrl,
		Width:     imageInfo.Width,
		Height:    imageInfo.Height,
		Extension: imageInfo.Extension,
		Folder:    "images",
		CloudName: common.KeyCompS3,
	}

	if err := biz.imageRepo.AddNewImage(ctx, &data); err != nil {
		urlParts := strings.Split(fileUrl, "/images/")
		imageId := urlParts[len(urlParts)-1]
		fileKey := fmt.Sprintf("images/%s", imageId)
		if err := biz.s3Client.DeleteObject(ctx, fileKey); err != nil {
			fmt.Printf("Failed to delete file from S3: %v\n", err)
		}
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateImage.Error()).
			WithDebug(err.Error())
	}

	return &data, nil
}
