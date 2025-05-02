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
	GetAudiosByNoteIdInt64(ctx context.Context, noteId int64) ([]entity.Audio, error)
	DeleteAudio(ctx context.Context, id int) error
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
			CreatedAt:    audio.CreatedAt.String(),
			UpdatedAt:    audio.UpdatedAt.String(),
		},
	}, nil
}

func (s *grpcService) GetAudiosByNoteId(ctx context.Context, req *pb.GetAudiosByNoteIdReq) (*pb.PublicAudioListResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	audios, err := s.business.GetAudiosByNoteIdInt64(ctx, req.NoteId)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	result := make([]*pb.PublicAudioInfo, len(audios))
	for i, audio := range audios {
		info := &pb.PublicAudioInfo{
			Id:        int64(audio.Id),
			NoteId:    audio.NoteID,
			FileUrl:   audio.FileURL,
			CreatedAt: audio.CreatedAt.String(),
			UpdatedAt: audio.UpdatedAt.String(),
		}

		if audio.TranscriptID != nil {
			info.TranscriptId = *audio.TranscriptID
		}
		if audio.SummaryID != nil {
			info.SummaryId = *audio.SummaryID
		}

		result[i] = info
	}

	return &pb.PublicAudioListResp{
		Audios: result,
	}, nil
}

func (s *grpcService) DeleteAudio(ctx context.Context, req *pb.DeleteAudioReq) (*pb.DeleteAudioResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	err := s.business.DeleteAudio(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.DeleteAudioResp{
		Success: true,
	}, nil
}
