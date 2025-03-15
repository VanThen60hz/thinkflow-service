package Audio

type Business interface {
	// GetAudioById(ctx context.Context, id int) (*entity.Image, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

// func (s *grpcService) GetImageById(ctx context.Context, req *pb.GetImageByIdReq) (*pb.PublicImageInfoResp, error) {
// 	image, err := s.business.GetAudioById(ctx, int(req.Id))
// 	if err != nil {
// 		return nil, core.ErrInternalServerError.WithError(err.Error())
// 	}

// 	return &pb.PublicImageInfoResp{
// 		Image: &pb.PublicImageInfo{
// 			Id:        int64(image.Id),
// 			Url:       image.Url,
// 			Width:     int64(image.Width),
// 			Height:    int64(image.Height),
// 			Extension: image.Extension,
// 		},
// 	}, nil
// }
