package api

import (
	"fmt"
	"net/http"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/services/audio/entity"
	"thinkflow-service/services/upload"

	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) CreateAudioHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		// Get file from request
		file, err := c.FormFile("file")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError("file is required"))
			return
		}

		// Process audio to get metadata
		processor := upload.NewMediaProcessor()
		audioInfo, err := processor.ProcessAudio(file)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		// Set requester to context first to get user ID
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		// Generate unique filename using user ID and timestamp
		uid, _ := core.FromBase58(requester.GetSubject())
		userId := int(uid.GetLocalID())
		fileName := fmt.Sprintf("%d_%d%s", userId, time.Now().UnixNano(), audioInfo.Format)

		// Upload to S3
		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		uploader := upload.NewS3Uploader(s3Component)
		_, err = uploader.UploadFile(ctx, file, "audios")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		data := entity.AudioDataCreation{
			Url:        fileName, // Store just the filename
			Format:     audioInfo.Format,
			Duration:   audioInfo.Duration,
			UploadedAt: audioInfo.UploadedAt,
		}

		if err := api.business.CreateNewAudio(ctx, &data); err != nil {
			fileKey := fmt.Sprintf("audios/%s", fileName)
			if delErr := uploader.DeleteFile(ctx, fileKey); delErr != nil {
				fmt.Printf("Failed to delete file after db error: %v\n", delErr)
			}
			common.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}
