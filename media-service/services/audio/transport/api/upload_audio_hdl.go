package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) UploadAudioHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		noteId, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError("file is required"))
			return
		}

		tempFile := fmt.Sprintf("./tmp/%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
		if err := c.SaveUploadedFile(file, tempFile); err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.
				WithError("cannot save uploaded file").
				WithDebug(err.Error()))
			return
		}
		defer os.Remove(tempFile)

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		data, err := api.business.UploadAudio(ctx, tempFile, file, int64(noteId.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}
