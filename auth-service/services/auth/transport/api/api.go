package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const oauthStateString = "random-state-string"

type Business interface {
	Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error)
	Register(ctx context.Context, data *entity.AuthRegister) error
	ForgotPassword(ctx context.Context, data *entity.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, data *entity.ResetPasswordRequest) error
	LoginOrRegisterWithGoogle(ctx context.Context, userInfo *entity.OAuthGoogleUserInfo) (*entity.TokenResponse, error)
	LoginOrRegisterWithFacebook(ctx context.Context, userInfo *entity.OAuthFacebookUserInfo) (*entity.TokenResponse, error)
	Logout(ctx context.Context, accessToken string) error
	VerifyEmail(ctx context.Context, data *entity.EmailVerificationRequest) error
	ResendVerificationOTP(ctx context.Context, data *entity.ResendOTPRequest) error
}

type api struct {
	serviceCtx sctx.ServiceContext
	business   Business
}

func NewAPI(serviceCtx sctx.ServiceContext, business Business) *api {
	return &api{serviceCtx: serviceCtx, business: business}
}

func (api *api) LoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.AuthEmailPassword

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		response, err := api.business.Login(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "accessToken",
			Value:    response.AccessToken.Token,
			Path:     "/",
			Domain:   "", 
			MaxAge:   604800,
			HttpOnly: true,
			Secure:   common.IsHTTPS(c),
			SameSite: http.SameSiteNoneMode, 
		})

		c.JSON(http.StatusOK, core.ResponseData("Login successfully"))
	}
}

func (api *api) RegisterHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.AuthRegister

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := api.business.Register(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (api *api) ForgotPasswordHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.ForgotPasswordRequest

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := api.business.ForgotPassword(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (api *api) ResetPasswordHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.ResetPasswordRequest

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := api.business.ResetPassword(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (api *api) GoogleLoginHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		url := common.AppOAuth2Config.GoogleConfig.AuthCodeURL(oauthStateString)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func (api *api) GoogleCallbackHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		if c.Query("state") != oauthStateString {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid oauth state"})
			return
		}

		token, err := common.AppOAuth2Config.GoogleConfig.Exchange(context.Background(), c.Query("code"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "code exchange failed"})
			return
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed getting user info"})
			return
		}
		defer resp.Body.Close()

		var userInfo entity.OAuthGoogleUserInfo
		if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed decoding user info"})
			return
		}

		tokenResponse, err := api.business.LoginOrRegisterWithGoogle(c.Request.Context(), &userInfo)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "accessToken",
			Value:    tokenResponse.AccessToken.Token,
			Path:     "/",
			Domain:   "", 
			MaxAge:   604800,
			HttpOnly: true,
			Secure:   common.IsHTTPS(c),
			SameSite: http.SameSiteNoneMode, 
		})

		c.JSON(http.StatusOK, core.ResponseData("Login successfully"))
	}
}

func (api *api) FacebookLoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		opts := []oauth2.AuthCodeOption{
			oauth2.SetAuthURLParam("auth_type", "rerequest"),
		}
		url := common.AppOAuth2Config.FacebookConfig.AuthCodeURL(oauthStateString, opts...)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

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
			common.WriteErrorResponse(c, err)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "accessToken",
			Value:    tokenResponse.AccessToken.Token,
			Path:     "/",
			Domain:   "", 
			MaxAge:   604800,
			HttpOnly: true,
			Secure:   common.IsHTTPS(c),
			SameSite: http.SameSiteNoneMode, 
		})
		

		c.JSON(http.StatusOK, core.ResponseData("Login successfully"))
	}
}

func (api *api) LogoutHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := c.Cookie("accessToken")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrUnauthorized.WithError("missing access token in cookie"))
			return
		}

		if token == "" {
			common.WriteErrorResponse(c, core.ErrUnauthorized.WithError("empty access token"))
			return
		}

		err = api.business.Logout(c.Request.Context(), token)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "accessToken",
			Value:    "",
			Path:     "/",
			Domain:   "", 
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   common.IsHTTPS(c),
			SameSite: http.SameSiteNoneMode, 
		})
		

		c.JSON(http.StatusOK, core.ResponseData("Logout successfully"))
	}
}

func (api *api) VerifyEmailHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.EmailVerificationRequest

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := api.business.VerifyEmail(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData("Email verification successful"))
	}
}

func (api *api) ResendVerificationOTPHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.ResendOTPRequest

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := api.business.ResendVerificationOTP(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData("OTP sent successfully"))
	}
}
