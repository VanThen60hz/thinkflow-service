package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (s *grpcService) IntrospectToken(ctx context.Context, req *pb.IntrospectReq) (*pb.IntrospectResp, error) {
	claims, err := s.business.IntrospectToken(ctx, req.AccessToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.IntrospectResp{
		Tid: claims.ID,
		Sub: claims.Subject,
	}, nil
}