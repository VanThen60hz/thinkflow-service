package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/user/dto"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) UpdateUserProfileHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dto.UpdateUserProfileRequest
		if err := c.ShouldBind(&req); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		updateData, err := req.ToUserDataUpdate()
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError("invalid avatar id"))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateUserProfile(ctx, updateData); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
