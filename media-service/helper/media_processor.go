package helper

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
	// Thư viện để xử lý MP3, cần cài đặt qua go get
)

type MediaProcessor struct{}

func NewMediaProcessor() *MediaProcessor {
	return &MediaProcessor{}
}

type ImageInfo struct {
	Width     int64
	Height    int64
	Extension string
}

type AudioInfo struct {
	Format     string
	UploadedAt string
}

type AttachmentInfo struct {
	FileName   string
	Extension  string
	SizeBytes  int64
	UploadedAt string
}

func (p *MediaProcessor) ProcessImage(file *multipart.FileHeader) (*ImageInfo, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, fmt.Errorf("cannot read file: %v", err)
	}

	img, _, err := image.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("cannot decode image: %v", err)
	}

	bounds := img.Bounds()
	ext := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

	return &ImageInfo{
		Width:     int64(bounds.Dx()),
		Height:    int64(bounds.Dy()),
		Extension: ext,
	}, nil
}

func (p *MediaProcessor) ProcessAudio(file *multipart.FileHeader) (*AudioInfo, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, fmt.Errorf("cannot read file: %v", err)
	}

	format := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

	return &AudioInfo{
		Format:     format,
		UploadedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

func (p *MediaProcessor) ProcessAttachment(file *multipart.FileHeader) (*AttachmentInfo, error) {
	ext := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

	return &AttachmentInfo{
		FileName:   file.Filename,
		Extension:  ext,
		SizeBytes:  file.Size,
		UploadedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
