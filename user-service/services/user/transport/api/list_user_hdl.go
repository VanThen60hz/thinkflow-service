package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ListUserHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.UserFilter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		rp.Paging.Process()

		note, err := api.business.ListUsers(ctx, &rp.UserFilter, &rp.Paging)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for i := range note {
			note[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(note, rp.Paging, rp.UserFilter))
	}
}
