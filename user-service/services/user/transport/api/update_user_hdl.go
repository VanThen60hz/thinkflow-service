package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) UpdateUserHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := core.FromBase58(c.Param("user-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError("invalid user id"))
			return
		}

		var updateData entity.UserDataUpdate
		if err := c.ShouldBindJSON(&updateData); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := updateData.Validate(); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateUser(ctx, int(userId.GetLocalID()), &updateData); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData("User updated successfully"))
	}
}
