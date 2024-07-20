package service

import (
	"fmt"

	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/domain/repository"
	"github.com/hrm1810884/works-hai-backend/infrastructure/scripts"
)

type GenerateDrawingService struct {
	userRepository    repository.UserRepository
	drawingRepository repository.DrawingRepository
}

func NewGenerateDrawingService(userRepository repository.UserRepository, drawingRepository repository.DrawingRepository) (*GenerateDrawingService, error) {
	return &GenerateDrawingService{userRepository, drawingRepository}, nil
}

func (s *GenerateDrawingService) GenerateDrawing(aiPosition *user.Position) ([]byte, error) {
	drawingData := map[string][]byte{}

	positions := map[string]*user.Position{
		"top":    aiPosition.GetTop(),
		"bottom": aiPosition.GetBottom(),
		"right":  aiPosition.GetRight(),
		"left":   aiPosition.GetLeft(),
	}

	for direction, pos := range positions {
		if pos != nil {
			posData, err := s.userRepository.FindByPos(*pos)
			if err != nil {
				return nil, fmt.Errorf("failed to get %s data: %w", direction, err)
			}
			drawing, err := s.drawingRepository.DownloadDrawing(posData.GetUrl())
			if err != nil {
				return nil, fmt.Errorf("failed to download %s drawing: %w", direction, err)
			}
			drawingData[direction] = drawing
		} else {
			drawingData[direction] = nil
		}
	}

	data, err := scripts.GenerateAIDrawing(drawingData)
	if err != nil {
		return nil, fmt.Errorf("failed to generate drawing: %w", err)
	}

	return data, nil
}
