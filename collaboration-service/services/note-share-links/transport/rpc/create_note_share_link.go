package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (s *grpcService) CreateNoteShareLink(ctx context.Context, req *pb.CreateNoteShareLinkRequest) (*pb.CreateNoteShareLinkResponse, error) {
	creation := toEntityNoteShareLinkCreation(req.GetShareLink())
	link, err := s.business.CreateNoteShareLink(ctx, creation)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.CreateNoteShareLinkResponse{
		ShareLink: toPBNoteShareLink(link),
	}, nil
}
