package api

import (
	"fmt"
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/media/upload"

	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) DeleteImageHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("image-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		// Get image data to get URL for deletion
		image, err := api.business.GetImageById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		// Delete from database first
		if err := api.business.DeleteImage(ctx, int(uid.GetLocalID())); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		// Delete from S3
		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		uploader := upload.NewS3Uploader(s3Component)
		fileKey := fmt.Sprintf("images/%s", image.Url)
		if err := uploader.DeleteFile(ctx, fileKey); err != nil {
			// Log error but don't return it to client since the database record is already deleted
			fmt.Printf("Failed to delete file from S3: %v\n", err)
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (api *api) DeleteAudioHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("audio-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		// Get audio data to get URL for deletion
		audio, err := api.business.GetAudioById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.DeleteAudio(ctx, int(uid.GetLocalID())); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		uploader := upload.NewS3Uploader(s3Component)
		fileKey := fmt.Sprintf("audios/%s", audio.Url)
		if err := uploader.DeleteFile(ctx, fileKey); err != nil {
			fmt.Printf("Failed to delete file from S3: %v\n", err)
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
