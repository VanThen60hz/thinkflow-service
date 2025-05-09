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
	grpcNoteAddress   string // for client make grpc client connection
	grpcImageAddress  string // for client make grpc client connection
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
		3200,
		"gRPC Port. Default: 3200",
	)

	flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3201",
		"gRPC server address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcAuthAddress,
		"grpc-auth-address",
		"localhost:3101",
		"gRPC auth server address. Default: localhost:3101",
	)

	flag.StringVar(
		&c.grpcNoteAddress,
		"grpc-note-address",
		"localhost:3301",
		"gRPC note server address. Default: localhost:3301",
	)

	flag.StringVar(
		&c.grpcImageAddress,
		"grpc-media-address",
		"localhost:3401",
		"gRPC image server address. Default: localhost:3401",
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

func (c *config) GetGRPCNoteServiceAddress() string {
	return c.grpcNoteAddress
}

func (c *config) GetGRPCImageServiceAddress() string {
	return c.grpcImageAddress
}
