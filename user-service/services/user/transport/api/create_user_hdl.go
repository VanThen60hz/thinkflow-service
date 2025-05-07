package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/user/business"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) CreateUserHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var createData business.CreateUserData
		if err := c.ShouldBindJSON(&createData); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		user, err := api.business.CreateUserByAdmin(ctx, &createData)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		user.Mask()

		c.JSON(http.StatusOK, core.SuccessResponse(user, nil, nil))
	}
}
