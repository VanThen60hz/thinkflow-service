package rpc

import (
	"context"
	"encoding/json"
	"errors"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/mindmap/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	GetMindmapById(ctx context.Context, id int) (*entity.Mindmap, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) GetMindmapById(ctx context.Context, req *pb.GetMindmapByIdReq) (*pb.PublicMindmapInfoResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	mindmap, err := s.business.GetMindmapById(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	mindmapData := ""
	if mindmap.MindmapData != nil {
		jsonData, err := json.Marshal(mindmap.MindmapData)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError("failed to marshal mindmap data: " + err.Error())
		}
		mindmapData = string(jsonData)
	}

	return &pb.PublicMindmapInfoResp{
		Mindmap: &pb.PublicMindmapInfo{
			Id:          int64(mindmap.Id),
			MindmapData: mindmapData,
		},
	}, nil
}
