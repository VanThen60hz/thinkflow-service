package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"thinkflow-service/common"
	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) UploadAttachmentHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		noteIDStr := c.PostForm("note-id")
		noteID, err := core.FromBase58(noteIDStr)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		tempFile := fmt.Sprintf("./tmp/%s", file.Filename)
		if err := c.SaveUploadedFile(file, tempFile); err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}
		defer os.Remove(tempFile)

		data, err := os.ReadFile(tempFile)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		s3FileName := fmt.Sprintf("attachments/%s", file.Filename)
		fileUrl, err := s3Component.UploadFileData(ctx, data, s3FileName)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		ext := filepath.Ext(file.Filename)

		attachment := &entity.AttachmentCreation{
			SQLModel:  core.SQLModel{},
			NoteID:    int64(noteID.GetLocalID()),
			FileURL:   fileUrl,
			FileName:  file.Filename,
			Extension: ext,
			SizeBytes: file.Size,
			CloudName: "s3",
		}

		if err := api.business.CreateAttachment(ctx, attachment); err != nil {
			if delErr := s3Component.DeleteObject(ctx, s3FileName); delErr != nil {
				fmt.Printf("Failed to delete file from S3: %v\n", delErr)
			}
			common.WriteErrorResponse(c, err)
			return
		}

		attachment.Mask()
		c.JSON(http.StatusOK, core.ResponseData(attachment.FakeId))
	}
}

func (api *api) GetAttachmentHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("attachment-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		attachment, err := api.business.GetAttachment(ctx, int64(id.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		attachment.Mask()
		c.JSON(http.StatusOK, core.ResponseData(attachment))
	}
}

func (api *api) GetAttachmentsByNoteIDHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		noteID, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		attachments, err := api.business.GetAttachmentsByNoteID(ctx, int64(noteID.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		for i := range attachments {
			attachments[i].Mask()
		}

		c.JSON(http.StatusOK, core.ResponseData(attachments))
	}
}

func (api *api) DeleteAttachmentHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("attachment-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		attachment, err := api.business.GetAttachment(ctx, int64(id.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		if err := api.business.DeleteAttachment(ctx, int64(id.GetLocalID())); err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		s3Component := api.serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
		urlParts := strings.Split(attachment.FileURL, "/attachments/")
		attachmentId := urlParts[len(urlParts)-1]
		fileKey := fmt.Sprintf("attachments/%s", attachmentId)
		if err := s3Component.DeleteObject(ctx, fileKey); err != nil {
			fmt.Printf("Failed to delete file from S3: %v\n", err)
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (api *api) UpdateAttachmentHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("attachment-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var data entity.Attachment
		if err := c.ShouldBindJSON(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateAttachment(ctx, int64(id.GetLocalID()), &data); err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
