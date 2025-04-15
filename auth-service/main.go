package main

import (
	"os"

	"thinkflow-service/cmd"
	"thinkflow-service/common"
)

func main() {
	common.InitOAuth2Configs()
	common.InitOauthStateString(os.Getenv("APP_NAME"))

	cmd.Execute()
}
