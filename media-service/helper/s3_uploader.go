package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/VanThen60hz/service-context/component/s3c"
)

type S3Uploader struct {
	s3Component *s3c.S3Component
}

func NewS3Uploader(s3Component *s3c.S3Component) *S3Uploader {
	return &S3Uploader{s3Component: s3Component}
}

func (u *S3Uploader) UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(file.Filename)

	fileName := fmt.Sprintf("%d%s", file.Size, ext)

	url, err := u.s3Component.Upload(ctx, fileName, folder)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	return url, nil
}

func (u *S3Uploader) DeleteFile(ctx context.Context, fileKey string) error {
	if err := u.s3Component.DeleteObject(ctx, fileKey); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	return nil
}
