package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (s *grpcService) GetNoteShareLinkByID(ctx context.Context, req *pb.GetNoteShareLinkByIDRequest) (*pb.GetNoteShareLinkResponse, error) {
	link, err := s.business.GetNoteShareLinkByID(ctx, req.GetId())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.GetNoteShareLinkResponse{
		ShareLink: toPBNoteShareLink(link),
	}, nil
}
