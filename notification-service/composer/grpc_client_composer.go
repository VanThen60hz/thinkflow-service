package composer

import (
	"context"
	"log"

	"thinkflow-service/common"
	"thinkflow-service/proto/pb"

	sctx "github.com/VanThen60hz/service-context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	grpcAuthClient pb.AuthServiceClient
}

func (ac *authClient) IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error) {
	resp, err := ac.grpcAuthClient.IntrospectToken(ctx, &pb.IntrospectReq{AccessToken: accessToken})
	if err != nil {
		return "", "", err
	}

	return resp.Sub, resp.Tid, nil
}

func ComposeAuthRPCClient(serviceCtx sctx.ServiceContext) pb.AuthServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(configComp.GetGRPCAuthServerAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewAuthServiceClient(clientConn)
}

func ComposeUserRPCClient(serviceCtx sctx.ServiceContext) pb.UserServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(configComp.GetGRPCUserServerAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewUserServiceClient(clientConn)
}
