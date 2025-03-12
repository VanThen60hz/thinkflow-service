package Image

import (
	"context"
	"fmt"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/media/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	GetImageById(ctx context.Context, id int) (*entity.Image, error)
	GetImagesByIds(ctx context.Context, ids []int) ([]entity.Image, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) GetImageById(ctx context.Context, req *pb.GetImageByIdReq) (*pb.PublicImageInfoResp, error) {
	image, err := s.business.GetImageById(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.PublicImageInfoResp{
		Image: &pb.PublicImageInfo{
			Id:  int64(image.Id),
			Url: image.Url,
		},
	}, nil
}

func (s *grpcService) GetImagesByIds(ctx context.Context, req *pb.GetImagesByIdsReq) (*pb.PublicImagesInfoResp, error) {
	imageIDs := make([]int, len(req.Ids))

	for i := range imageIDs {
		imageIDs[i] = int(req.Ids[i])
	}

	fmt.Println("ImageIDs", imageIDs)

	images, err := s.business.GetImagesByIds(ctx, imageIDs)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	publicImageInfo := make([]*pb.PublicImageInfo, len(images))

	for i := range images {
		publicImageInfo[i] = &pb.PublicImageInfo{
			Id:  int64(images[i].Id),
			Url: images[i].Url,
		}
	}

	return &pb.PublicImagesInfoResp{Images: publicImageInfo}, nil
}
