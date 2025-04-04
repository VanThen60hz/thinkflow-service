package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/services/image/entity"
	"thinkflow-service/services/upload"

	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) CreateImageHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		// Get file from request
		file, err := c.FormFile("file")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError("file is required"))
			return
		}

		// Process image to get metadata
		processor := upload.NewMediaProcessor()
		imageInfo, err := processor.ProcessImage(file)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		tempFile := fmt.Sprintf("./tmp/%d%s", time.Now().UnixNano(), imageInfo.Extension)
		if err := c.SaveUploadedFile(file, tempFile); err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}
		defer os.Remove(tempFile)

		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		fileUrl, err := s3Component.Upload(ctx, tempFile, "images")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		data := entity.ImageDataCreation{
			Url:       fileUrl,
			Width:     imageInfo.Width,
			Height:    imageInfo.Height,
			Extension: imageInfo.Extension,
			Folder:    "images",
			CloudName: common.KeyCompS3,
		}

		if err := api.business.CreateNewImage(ctx, &data); err != nil {
			urlParts := strings.Split(fileUrl, "/images/")
			imageId := urlParts[len(urlParts)-1]
			fileKey := fmt.Sprintf("images/%s", imageId)
			if err := s3Component.DeleteObject(ctx, fileKey); err != nil {
				fmt.Printf("Failed to delete file from S3: %v\n", err)
			}
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}
