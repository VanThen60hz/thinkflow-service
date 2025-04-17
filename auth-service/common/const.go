package common

import "github.com/VanThen60hz/service-context/core"

const (
	KeyCompMySQL = "mysql"
	KeyCompGIN   = "gin"
	KeyCompJWT   = "jwt"
	KeyCompConf  = "config"
	KeyCompRedis = "redis"
	KeyCompEmail = "email"
	KeyCompOAuth = "oauth"

	MaskTypeUser = 1
	MaskTypeNote = 2

	EmailVerifyOTPSubject     = "ThinkFlow - Email Verification OTP"
	EmailResetPasswordSubject = "ThinkFlow - Password Reset OTP"

	RequesterKey = string(core.KeyRequester)
)
