package service

import (
	"fmt"

	"github.com/hrm1810884/works-hai-backend/domain/repository"
)

type GetSignedUrlService struct {
	repository repository.DrawingRepository
}

func NewGetSignedUrlService(drawingRepository repository.DrawingRepository) (*GetSignedUrlService, error) {
	return &GetSignedUrlService{repository: drawingRepository}, nil
}

func (s *GetSignedUrlService) GetSignedUrl(drawingName string, method string) (string, error) {
	presignedUrl, err := s.repository.GenerateSignedUrl(drawingName, 15, method)
	if err != nil {
		return "", fmt.Errorf("error generating url: %w", err)
	}

	return presignedUrl, nil
}
