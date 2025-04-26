package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetUserProfileHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		user, err := api.business.GetUserProfile(ctx)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		user.Mask()

		c.JSON(http.StatusOK, core.ResponseData(user))
	}
}
