package usecase

import (
	"context"
	"fmt"

	"github.com/hrm1810884/works-hai-backend/entity"
	"github.com/hrm1810884/works-hai-backend/usecase/service"
)

type IQuadSignedUrls interface {
	GenerateQuadSignedUrls(ctx context.Context, currentPosition entity.IPosition) (map[string]string, error)
}

type QuadSignedUrlsUsecase struct {
	QuadImages entity.QuadImagesEntity
}

func NewQuadSignedUrlsUsecase(ctx context.Context, currentPosition entity.IPosition) IQuadSignedUrls {
	quadImages := entity.NewQuadImages(ctx, currentPosition)
	return &QuadSignedUrlsUsecase{
		QuadImages: quadImages,
	}
}

func (u *QuadSignedUrlsUsecase) GenerateQuadSignedUrls(ctx context.Context, currentPosition entity.IPosition) (map[string]string, error) {
	s, err := service.NewGetSignedUrlService(ctx)
	if err != nil {
		return nil, err
	}

	humanDrawingFileName := fmt.Sprintf("%v_%v.png", currentPosition.GetX(), currentPosition.GetY())
	humanSignedUrl, err := s.GetSignedUrl(humanDrawingFileName, "PUT")
	if err != nil {
		return nil, err
	}

	signedUrls := map[string]string{
		"Human": humanSignedUrl,
	}

	for pos, cfg := range u.QuadImages.Config {
		if cfg.IsDrawn {
			url, err := s.GetSignedUrl(cfg.ResourceName, "GET")
			if err != nil {
				return nil, err
			}
			signedUrls[pos] = url
		} else {
			signedUrls[pos] = ""
		}
	}

	return signedUrls, nil
}
