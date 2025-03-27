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
			CreatedAt:   Summary.CreatedAt.String(),
			UpdatedAt:   Summary.UpdatedAt.String(),
		},
	}, nil
}
