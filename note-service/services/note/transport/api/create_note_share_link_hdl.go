package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type CreateNoteShareLinkRequest struct {
	Permission  string `json:"permission" binding:"required,oneof=read write"`
	ExpiryTimeM *int   `json:"expiry_time"` // số phút, có thể null
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
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var req CreateNoteShareLinkRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		expiresAt := calculateExpiresAt(req.ExpiryTimeM)

		link, err := api.business.CreateNoteShareLink(ctx, int64(noteId.GetLocalID()), req.Permission, expiresAt)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		clientURL := os.Getenv("CLIENT_URL")
		clientURL = strings.TrimRight(clientURL, "/")
		shareURL := fmt.Sprintf("%s/share/%s", clientURL, link.Token)

		c.JSON(http.StatusOK, core.ResponseData(ShareLinkResponse{URL: shareURL}))
	}
}

func calculateExpiresAt(minutes *int) *time.Time {
	if minutes == nil {
		return nil
	}
	t := time.Now().Add(time.Duration(*minutes) * time.Minute)
	return &t
}
