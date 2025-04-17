package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/helper"
	"thinkflow-service/services/audio/entity"

	// Thêm import package upload
	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) CreateAudioHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		noteId, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError("file is required"))
			return
		}

		processor := helper.NewMediaProcessor()
		audioInfo, err := processor.ProcessAudio(file)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		tempFile := fmt.Sprintf("./tmp/%d.%s", time.Now().UnixNano(), audioInfo.Format)
		if err := c.SaveUploadedFile(file, tempFile); err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}
		defer os.Remove(tempFile)

		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		fileUrl, err := s3Component.Upload(ctx, tempFile, "audios")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		data := entity.AudioDataCreation{
			NoteID:  int64(noteId.GetLocalID()),
			FileURL: fileUrl,
			Format:  audioInfo.Format,
		}

		if err := api.business.CreateNewAudio(ctx, &data); err != nil {
			urlParts := strings.Split(fileUrl, "/audios/")
			AudioId := urlParts[len(urlParts)-1]
			fileKey := fmt.Sprintf("Audios/%s", AudioId)
			if err := s3Component.DeleteObject(ctx, fileKey); err != nil {
				fmt.Printf("Failed to delete file from S3: %v\n", err)
			}
			common.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}
