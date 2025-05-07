package rpc

import (
	"context"

	"thinkflow-service/proto/pb"
)

func (s *grpcService) CountNotes(ctx context.Context, req *pb.CountNotesReq) (*pb.CountNotesResp, error) {
	totalNotes, err := s.business.CountNotes(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CountNotesResp{
		TotalNotes: totalNotes,
	}, nil
}
