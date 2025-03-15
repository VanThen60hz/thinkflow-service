package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) UpdateImageHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("image-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var data entity.ImageDataUpdate

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateImage(ctx, int(uid.GetLocalID()), &data); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
