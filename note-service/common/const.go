package common

import "github.com/VanThen60hz/service-context/core"

type contextKey string

const (
	KeyCompMySQL = "mysql"
	KeyCompGIN   = "gin"
	KeyCompJWT   = "jwt"
	KeyCompConf  = "config"

	MaskTypeUser = 1
	MaskTypeNote = 2

	RequesterKey contextKey = core.KeyRequester
)
