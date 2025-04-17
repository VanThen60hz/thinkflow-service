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

func (api *api) CreateImageHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError("file is required"))
			return
		}

		tempFile := fmt.Sprintf("./tmp/%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

		if err := c.SaveUploadedFile(file, tempFile); err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.
				WithError("cannot save uploaded file").
				WithDebug(err.Error()))
			return
		}
		defer os.Remove(tempFile)

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		data, err := api.business.UploadImage(ctx, tempFile, file)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}
