package api

import (
	"context"
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type Business interface {
	GetUserProfile(ctx context.Context) (*entity.User, error)
	UpdateUserProfile(ctx context.Context, data *entity.UserDataUpdate) error
}

type api struct {
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}

func (api *api) GetUserProfileHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		user, err := api.business.GetUserProfile(ctx)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		user.Mask()

		c.JSON(http.StatusOK, core.ResponseData(user))
	}
}

type updateUserProfileRequest struct {
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Gender    *entity.Gender `json:"gender,omitempty"`
	AvatarId  *string        `json:"avatar_id,omitempty"`
}

func (api *api) UpdateUserProfileHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req updateUserProfileRequest
		if err := c.ShouldBind(&req); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		updateData := &entity.UserDataUpdate{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Phone:     req.Phone,
			Gender:    req.Gender,
		}

		// Handle avatar update if provided
		if req.AvatarId != nil {
			avatarUid, err := core.FromBase58(*req.AvatarId)
			if err != nil {
				common.WriteErrorResponse(c, core.ErrBadRequest.WithError("invalid avatar id"))
				return
			}
			avatarId := int(avatarUid.GetLocalID())
			updateData.Avatar = &avatarId
		}

		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateUserProfile(ctx, updateData); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
