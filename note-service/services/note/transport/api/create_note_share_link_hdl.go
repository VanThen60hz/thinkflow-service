package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type CreateNoteShareLinkRequest struct {
	Permission string `json:"permission" binding:"required,oneof=read write"`
}

type ShareLinkResponse struct {
	URL string `json:"url"`
}

func (api *api) CreateNoteShareLinkHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		noteIdStr := c.Param("note-id")
		noteId, err := core.FromBase58(noteIdStr)
		if err != nil {
			return
		}

		var req CreateNoteShareLinkRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			return
		}

		link, err := api.business.CreateNoteShareLink(ctx, int64(noteId.GetLocalID()), req.Permission)
		if err != nil {
			return
		}

		clientURL := os.Getenv("CLIENT_URL")
		clientURL = strings.TrimRight(clientURL, "/")
		shareURL := fmt.Sprintf("%s/share/%s", clientURL, link.Token)
		response := ShareLinkResponse{
			URL: shareURL,
		}

		c.JSON(http.StatusOK, core.ResponseData(response))
	}
}
