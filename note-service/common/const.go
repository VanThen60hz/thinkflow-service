package common

import "github.com/VanThen60hz/service-context/core"

type contextKey string

const (
	KeyCompMySQL = "mysql"
	KeyCompGIN   = "gin"
	KeyCompJWT   = "jwt"
	KeyCompConf  = "config"

	MaskTypeUser          = 1
	MaskTypeNote          = 2
	MaskTypeImage         = 3
	MaskTypeAudio         = 4
	MaskTypeTranscript    = 5
	MaskTypeSummary       = 6
	MaskTypeMindmap       = 7
	MaskTypeText          = 8
	MaskTypeCollaboration = 9

	RequesterKey contextKey = core.KeyRequester
)
