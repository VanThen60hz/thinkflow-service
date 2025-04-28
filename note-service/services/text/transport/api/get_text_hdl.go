package api

import (
	"net/http"
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetTextHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		tid, err := core.FromBase58(c.Param("text-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		data, err := api.business.GetTextById(ctx, int(tid.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
