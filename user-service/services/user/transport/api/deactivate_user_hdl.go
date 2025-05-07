package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) DeactivateUserHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := core.FromBase58(c.Param("user-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError("invalid user id"))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		result, err := api.business.DeactivateUser(ctx, int(userId.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.SuccessResponse(result, nil, nil))
	}
}
