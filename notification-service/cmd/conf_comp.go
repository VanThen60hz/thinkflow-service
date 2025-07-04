package cmd

import (
	"flag"

	"thinkflow-service/common"

	sctx "github.com/VanThen60hz/service-context"
)

type config struct {
	grpcPort          int    // for server port listening
	grpcServerAddress string // for client make grpc client connection
	grpcAuthAddress   string // for client make grpc client connection
	grpcUserAddress   string // for client make grpc client connection
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
		3700,
		"gRPC Port. Default: 3700",
	)

	flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3701",
		"gRPC server address. Default: localhost:3701",
	)

	flag.StringVar(
		&c.grpcAuthAddress,
		"grpc-auth-address",
		"localhost:3101",
		"gRPC auth server address. Default: localhost:3101",
	)

	flag.StringVar(
		&c.grpcUserAddress,
		"grpc-user-address",
		"localhost:3201",
		"gRPC auth server address. Default: localhost:3201",
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

func (c *config) GetGRPCServerAddress() string {
	return c.grpcServerAddress
}

func (c *config) GetGRPCAuthServerAddress() string {
	return c.grpcAuthAddress
}

func (c *config) GetGRPCUserServerAddress() string {
	return c.grpcUserAddress
}
