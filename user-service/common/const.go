package common

import "github.com/VanThen60hz/service-context/core"

const (
	KeyCompMySQL = "mysql"
	KeyCompGIN   = "gin"
	KeyCompJWT   = "jwt"
	KeyCompConf  = "config"

	MaskTypeUser  = 1
	MaskTypeNote  = 2
	MaskTypeImage = 3

	RequesterKey = string(core.KeyRequester)
)
