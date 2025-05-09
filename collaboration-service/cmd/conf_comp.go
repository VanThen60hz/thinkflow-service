package cmd

import (
	"flag"

	"thinkflow-service/common"

	sctx "github.com/VanThen60hz/service-context"
)

type config struct {
	grpcPort              int    // for server port listening
	grpcAuthServerAddress string // for client make grpc client connection
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ID() string {
	return common.KeyCompConf
}

func (c *config) InitFlags() {
	flag.IntVar(
		&c.grpcPort,
		"grpc-port",
		3300,
		"gRPC Port. Default: 3300",
	)

	flag.StringVar(
		&c.grpcAuthServerAddress,
		"grpc-auth-address",
		"localhost:3101",
		"gRPC server address. Default: localhost:3101",
	)
}

func (c *config) Activate(_ sctx.ServiceContext) error {
	return nil
}

func (c *config) Stop() error {
	return nil
}

func (c *config) GetGRPCPort() int {
	return c.grpcPort
}

func (c *config) GetGRPCAuthServerAddress() string {
	return c.grpcAuthServerAddress
}
