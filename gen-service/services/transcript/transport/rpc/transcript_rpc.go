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
	CreateTranscript(ctx context.Context, content string) (int, error)
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

func (s *grpcService) CreateTranscript(ctx context.Context, req *pb.CreateTranscriptReq) (*pb.NewTranscriptIdResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	id, err := s.business.CreateTranscript(ctx, req.Content)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.NewTranscriptIdResp{
		Id: int32(id),
	}, nil
}
