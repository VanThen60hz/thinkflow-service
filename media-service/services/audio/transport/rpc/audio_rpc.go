package rpc

import (
	"context"
	"errors"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	GetAudioById(ctx context.Context, id int) (*entity.Audio, error)
	GetAudiosByNoteId(ctx context.Context, noteId int) ([]entity.Audio, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) GetAudioById(ctx context.Context, req *pb.GetAudioByIdReq) (*pb.PublicAudioInfoResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	audio, err := s.business.GetAudioById(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.PublicAudioInfoResp{
		Audio: &pb.PublicAudioInfo{
			Id:           int64(audio.Id),
			NoteId:       audio.NoteID,
			FileUrl:      audio.FileURL,
			TranscriptId: *audio.TranscriptID,
			SummaryId:    *audio.SummaryID,
			MindmapId:    *audio.MindmapID,
			CreatedAt:    audio.CreatedAt.String(),
			UpdatedAt:    audio.UpdatedAt.String(),
		},
	}, nil
}

func (s *grpcService) GetAudiosByNoteId(ctx context.Context, req *pb.GetAudiosByNoteIdReq) (*pb.PublicAudioListResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	audios, err := s.business.GetAudiosByNoteId(ctx, int(req.NoteId))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	result := make([]*pb.PublicAudioInfo, len(audios))
	for i, audio := range audios {
		result[i] = &pb.PublicAudioInfo{
			Id:           int64(audio.Id),
			NoteId:       audio.NoteID,
			FileUrl:      audio.FileURL,
			TranscriptId: *audio.TranscriptID,
			SummaryId:    *audio.SummaryID,
			MindmapId:    *audio.MindmapID,
			CreatedAt:    audio.CreatedAt.String(),
			UpdatedAt:    audio.UpdatedAt.String(),
		}
	}

	return &pb.PublicAudioListResp{
		Audios: result,
	}, nil
}
