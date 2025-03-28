package cmd

import (
	"flag"

	"thinkflow-service/common"

	sctx "github.com/VanThen60hz/service-context"
)

type config struct {
	grpcPort               int    // for server port listening
	grpcServerAddress      string // for client make grpc client connection
	grpcAuthServerAddress  string // for client make grpc client connection
	grpcUserServiceAddress string // for client make grpc client connection
	grpcGenServiceAddress  string // for client make grpc client connection
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
		3400,
		"gRPC Port. Default: 3400",
	)

	flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3401",
		"gRPC server address. Default: localhost:3401",
	)

	flag.StringVar(
		&c.grpcAuthServerAddress,
		"grpc-auth-address",
		"localhost:3101",
		"gRPC server address. Default: localhost:3101",
	)

	flag.StringVar(
		&c.grpcUserServiceAddress,
		"grpc-user-address",
		"localhost:3201",
		"gRPC server address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcGenServiceAddress,
		"grpc-gen-address",
		"localhost:3501",
		"gRPC gen server address. Default: localhost:3501",
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

func (c *config) GetGRPCUserServiceAddress() string {
	return c.grpcUserServiceAddress
}

func (c *config) GetGRPCGenServiceAddress() string {
	return c.grpcGenServiceAddress
}
