package rpc

import (
	"context"
	"errors"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/transcript/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	GetTranscriptById(ctx context.Context, id int) (*entity.Transcript, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) GetTranscriptById(ctx context.Context, req *pb.GetTranscriptByIdReq) (*pb.PublicTranscriptInfoResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	transcript, err := s.business.GetTranscriptById(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.PublicTranscriptInfoResp{
		Transcript: &pb.PublicTranscriptInfo{
			Id:      int64(transcript.Id),
			Content: transcript.Content,
		},
	}, nil
}
