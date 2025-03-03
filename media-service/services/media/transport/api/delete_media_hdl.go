package api

import (
	"fmt"
	"net/http"
	"strings"

	"thinkflow-service/common"

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

		image, err := api.business.GetImageById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.DeleteImage(ctx, int(uid.GetLocalID())); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		urlParts := strings.Split(image.Url, "/images/")
		imageId := urlParts[len(urlParts)-1]
		fileKey := fmt.Sprintf("images/%s", imageId)
		if err := s3Component.DeleteObject(ctx, fileKey); err != nil {
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
		fileKey := fmt.Sprintf("audios/%s", audio.Url)
		if err := s3Component.DeleteObject(ctx, fileKey); err != nil {
			fmt.Printf("Failed to delete file from S3: %v\n", err)
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
