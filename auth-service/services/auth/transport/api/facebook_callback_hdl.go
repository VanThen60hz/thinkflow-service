package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) FacebookCallbackHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		if c.Query("state") != oauthStateString {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid oauth state"})
			return
		}

		token, err := common.AppOAuth2Config.FacebookConfig.Exchange(context.Background(), c.Query("code"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "code exchange failed"})
			return
		}

		resp, err := http.Get("https://graph.facebook.com/v20.0/me?fields=id,name,email&access_token=" + token.AccessToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed getting user info"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed reading response"})
			return
		}

		var userInfo entity.OAuthFacebookUserInfo
		if err := json.Unmarshal(body, &userInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed parsing user info"})
			return
		}

		tokenResponse, err := api.business.LoginOrRegisterWithFacebook(c.Request.Context(), &userInfo)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "accessToken",
			Value:    tokenResponse.AccessToken.Token,
			Path:     "/",
			Domain:   "",
			MaxAge:   604800,
			HttpOnly: true,
			Secure:   core.IsHTTPS(c),
			SameSite: http.SameSiteNoneMode,
		})

		c.JSON(http.StatusOK, core.ResponseData("Login successfully"))
	}
}
