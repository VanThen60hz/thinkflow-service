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

func (api *api) UploadAttachmentHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		noteIDStr := c.PostForm("note-id")
		noteID, err := core.FromBase58(noteIDStr)
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		tempFile := fmt.Sprintf("./tmp/%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
		if err := c.SaveUploadedFile(file, tempFile); err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}
		defer os.Remove(tempFile)

		attachment, err := api.business.UploadAttachment(ctx, tempFile, file, int64(noteID.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		attachment.Mask()
		c.JSON(http.StatusOK, core.ResponseData(attachment.FakeId))
	}
}
