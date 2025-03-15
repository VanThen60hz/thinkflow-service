package upload

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
	Duration   int64
	UploadedAt string
}

func (p *MediaProcessor) ProcessImage(file *multipart.FileHeader) (*ImageInfo, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	defer src.Close()

	// Read file content
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, fmt.Errorf("cannot read file: %v", err)
	}

	// Decode image
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

	format := strings.TrimPrefix(filepath.Ext(file.Filename), ".")
	duration := int64(0)

	return &AudioInfo{
		Format:     format,
		Duration:   duration,
		UploadedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
