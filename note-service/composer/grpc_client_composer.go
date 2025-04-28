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

// ComposeAuthRPCClient use only for middleware: get token info
func ComposeAuthRPCClient(serviceCtx sctx.ServiceContext) *authClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(configComp.GetGRPCAuthServerAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return &authClient{pb.NewAuthServiceClient(clientConn)}
}

func composeUserRPCClient(serviceCtx sctx.ServiceContext) pb.UserServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(configComp.GetGRPCUserServiceAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewUserServiceClient(clientConn)
}

func ComposeSummaryRPCClient(serviceCtx sctx.ServiceContext) pb.SummaryServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(configComp.GetGRPCGenServiceAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewSummaryServiceClient(clientConn)
}

func ComposeMindmapRPCClient(serviceCtx sctx.ServiceContext) pb.MindmapServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(configComp.GetGRPCGenServiceAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewMindmapServiceClient(clientConn)
}

func ComposeCollaborationRPCClient(serviceCtx sctx.ServiceContext) pb.CollaborationServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient(configComp.GetGRPCCollaborationServiceAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewCollaborationServiceClient(clientConn)
}