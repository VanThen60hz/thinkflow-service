package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/media/entity"

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
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		rp.Paging.Process()

		images, err := api.business.ListImages(c.Request.Context(), &rp.Filter, &rp.Paging)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		for i := range images {
			images[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(images, rp.Paging, rp.Filter))
	}
}

func (api *api) ListAudiosHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		rp.Paging.Process()

		audios, err := api.business.ListAudios(c.Request.Context(), &rp.Filter, &rp.Paging)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		for i := range audios {
			audios[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(audios, rp.Paging, rp.Filter))
	}
}
