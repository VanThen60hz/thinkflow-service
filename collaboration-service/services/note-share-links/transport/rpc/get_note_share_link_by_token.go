package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/pkg/errors"
)

func (s *grpcService) GetNoteShareLinkByToken(ctx context.Context, req *pb.GetNoteShareLinkByTokenRequest) (*pb.GetNoteShareLinkResponse, error) {
	link, err := s.business.GetNoteShareLinkByToken(ctx, req.GetToken())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.GetNoteShareLinkResponse{
		ShareLink: toPBNoteShareLink(link),
	}, nil
}
