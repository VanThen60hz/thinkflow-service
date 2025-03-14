package main

import (
	"os"

	"thinkflow-service/cmd"
)

func main() {
	os.Setenv("REDIS_ADDRESS", os.Getenv("REDIS_ADDRESS"))
	os.Setenv("EMAIL_USER", os.Getenv("EMAIL_USER"))
	os.Setenv("EMAIL_PASSWORD", os.Getenv("EMAIL_PASSWORD"))

	cmd.Execute()
}
