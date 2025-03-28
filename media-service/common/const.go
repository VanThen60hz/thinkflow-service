package common

import "github.com/VanThen60hz/service-context/core"

type contextKey string

const (
	KeyCompMySQL = "mysql"
	KeyCompGIN   = "gin"
	KeyCompJWT   = "jwt"
	KeyCompConf  = "config"
	KeyCompS3    = "s3"

	MaskTypeUser       = 1
	MaskTypeNote       = 2
	MaskTypeImage      = 3
	MaskTypeAudio      = 4
	MaskTypeTranscript = 5
	MaskTypeSummary    = 6
	MaskTypeMindmap    = 7

	RequesterKey contextKey = core.KeyRequester
)
