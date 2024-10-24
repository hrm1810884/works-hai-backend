package service

import (
	"fmt"

	repository "github.com/hrm1810884/works-hai-backend/domain/repository"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
)

type GetSignedUrlService struct {
	repository repository.DrawingRepository
}

func NewGetSignedUrlService(implDrawingRepository *impl_repository.ImplDrawingRepository) (*GetSignedUrlService, error) {
	drawingRepository := repository.DrawingRepository(implDrawingRepository)
	return &GetSignedUrlService{repository: drawingRepository}, nil
}

func (s *GetSignedUrlService) GetSignedUrl(drawingName string, method string) (string, error) {
	presignedUrl, err := s.repository.GenerateSignedUrl(drawingName, method)
	if err != nil {
		return "", fmt.Errorf("error generating url: %w", err)
	}

	return presignedUrl, nil
}
