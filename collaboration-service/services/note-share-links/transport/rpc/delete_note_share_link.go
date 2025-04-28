package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (s *grpcService) DeleteNoteShareLink(ctx context.Context, req *pb.DeleteNoteShareLinkRequest) (*pb.DeleteNoteShareLinkResponse, error) {
	if err := s.business.DeleteNoteShareLink(ctx, req.GetId()); err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.DeleteNoteShareLinkResponse{
		Success: true,
	}, nil
}
