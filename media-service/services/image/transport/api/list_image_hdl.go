package api

import (
	"net/http"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ListImagesHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		rp.Paging.Process()

		images, err := api.business.ListImages(c.Request.Context(), &rp.Filter, &rp.Paging)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for i := range images {
			images[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(images, rp.Paging, rp.Filter))
	}
}
