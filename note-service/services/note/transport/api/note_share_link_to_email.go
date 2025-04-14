package api

import (
	"fmt"
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type NoteShareLinkToEmailRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Permission  string `json:"permission" binding:"required,oneof=read write"`
	ExpiryTimeM *int   `json:"expiry_time"` // số phút, có thể null
}

func (api *api) NoteShareLinkToEmailHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		noteIdStr := c.Param("note-id")
		noteId, err := core.FromBase58(noteIdStr)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var req NoteShareLinkToEmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		expiresAt := calculateExpiresAt(req.ExpiryTimeM)

		fmt.Println("Note ID:", noteId.GetLocalID())
		fmt.Println("Email:", req.Email)
		fmt.Println("Permission:", req.Permission)
		fmt.Println("Expires At:", expiresAt)

		err = api.business.NoteShareLinkToEmail(ctx, int64(noteId.GetLocalID()), req.Email, req.Permission, expiresAt)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData("Share link sent to email successfully"))
	}
}
