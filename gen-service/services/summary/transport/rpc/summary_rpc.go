package rpc

import (
	"context"
	"errors"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/summary/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	GetSummaryById(ctx context.Context, id int) (*entity.Summary, error)
	CreateSummary(ctx context.Context, summaryText string) (int, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) GetSummaryById(ctx context.Context, req *pb.GetSummaryByIdReq) (*pb.PublicSummaryInfoResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	Summary, err := s.business.GetSummaryById(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.PublicSummaryInfoResp{
		Summary: &pb.PublicSummaryInfo{
			Id:          int64(Summary.Id),
			SummaryText: Summary.SummaryText,
		},
	}, nil
}

func (s *grpcService) CreateSummary(ctx context.Context, req *pb.CreateSummaryReq) (*pb.NewSummaryIdResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	id, err := s.business.CreateSummary(ctx, req.SummaryText)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.NewSummaryIdResp{
		Id: int32(id),
	}, nil
}
