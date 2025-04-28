package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (s *grpcService) UpdateNoteShareLink(ctx context.Context, req *pb.UpdateNoteShareLinkRequest) (*pb.UpdateNoteShareLinkResponse, error) {
	update := toEntityNoteShareLinkUpdate(req.GetShareLink())
	link, err := s.business.UpdateNoteShareLink(ctx, req.GetId(), update)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.UpdateNoteShareLinkResponse{
		ShareLink: toPBNoteShareLink(link),
	}, nil
}
