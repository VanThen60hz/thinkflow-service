package common

import "github.com/VanThen60hz/service-context/core"

const (
	KeyCompMySQL = "mysql"
	KeyCompGIN   = "gin"
	KeyCompJWT   = "jwt"
	KeyCompConf  = "config"
	KeyCompNats  = "nats"

	MaskTypeUser         = 1
	MaskTypeNote         = 2
	MaskTypeImage        = 3
	MaskTypeAudio        = 4
	MaskTypeTranscript   = 5
	MaskTypeSummary      = 6
	MaskTypeMindmap      = 7
	MaskTypeAttachment   = 8
	MaskTypeNotification = 9

	RequesterKey = string(core.KeyRequester)

	TimeFormat = "2006-01-02T15:04:05Z07:00"
)
