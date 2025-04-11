package main

import (
	"thinkflow-service/cmd"
	"thinkflow-service/common"
)

func main() {
	common.InitOAuth2Configs()

	cmd.Execute()
}
