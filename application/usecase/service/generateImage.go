package service

import (
	"errors"
	"fmt"

	user "github.com/hrm1810884/works-hai-backend/domain/entity/user"
	repository "github.com/hrm1810884/works-hai-backend/domain/repository"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
)

type DrawingService struct {
	userRepository    repository.UserRepository
	drawingRepository repository.DrawingRepository
}

func NewDrawingService(implUserRepository *impl_repository.ImplUserRepository, implDrawingRepository *impl_repository.ImplDrawingRepository) (*DrawingService, error) {
	var userRepository = repository.UserRepository(implUserRepository);
	var drawingRepository = repository.DrawingRepository(implDrawingRepository);

	return &DrawingService{userRepository, drawingRepository}, nil
}

func (s *DrawingService) GetQuadUrls(pivotPosition user.Position) (map[string]string, error) {
	quadUrls := map[string]string{}

	positions := map[string]*user.Position{
		"top":    pivotPosition.GetTop(),
		"bottom": pivotPosition.GetBottom(),
		"right":  pivotPosition.GetRight(),
		"left":   pivotPosition.GetLeft(),
	}

	for direction, pos := range positions {
		if pos != nil {
			posData, err := s.userRepository.FindByPos(*pos)
			if err != nil {
				return nil, fmt.Errorf("failed to get %s data: %w", direction, err)
			}
			quadUrls[direction] = posData.GetUrl()
		} else {
			quadUrls[direction] = ""
		}
	}

	if err := validateQuadMap(quadUrls); err != nil {
		return nil, err
	}

	return quadUrls, nil

}

func (s *DrawingService) GenerateDrawing(aiPosition *user.Position) ([]byte, error) {
	drawingData := map[string][]byte{}

	quadUrls, err := s.GetQuadUrls(*aiPosition)
	if err != nil {
		return nil, fmt.Errorf("failed to get quad urls: %w", err)
	}

	for direction, url := range quadUrls {
		if url != "" {
			drawing, err := s.drawingRepository.DownloadDrawing(url)
			if err != nil {
				return nil, fmt.Errorf("failed to download %s drawing: %w", direction, err)
			}
			drawingData[direction] = drawing
		} else {
			drawingData[direction] = nil
		}
	}

	if err := validateQuadMap(drawingData); err != nil {
		return nil, err
	}

	data, err := s.drawingRepository.GenerateAIDrawing(drawingData)
	if err != nil {
		return nil, fmt.Errorf("failed to generate drawing: %w", err)
	}

	return data, nil
}

func validateQuadMap(quadData interface{}) error {
	convertedData, err := convertToMapInterface(quadData)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := validateKey(convertedData); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := validateValue(convertedData); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}

func validateKey(quadData map[string]interface{}) error {
	for key := range quadData {
		if key != "top" && key != "bottom" && key != "right" && key != "left" {
			return fmt.Errorf("invalid key found: %s", key)
		}
	}

	return nil
}

func validateValue(quadData map[string]interface{}) error {
	allEmpty := true
	for _, val := range quadData {
		switch v := val.(type) {
		case string:
			if v != "" {
				allEmpty = false
			}
		case nil:
			// do nothing
		default:
			allEmpty = false
		}

		if !allEmpty {
			break
		}
	}

	if allEmpty {
		return fmt.Errorf("all directions have empty URLs")
	}

	return nil
}

func convertToMapInterface(originalMap interface{}) (map[string]interface{}, error) {
	convertedMap := make(map[string]interface{})
	switch m := originalMap.(type) {
	case map[string]string:
		for k, v := range m {
			convertedMap[k] = v
		}
	case map[string][]byte:
		for k, v := range m {
			convertedMap[k] = v
		}
	default:
		return nil, errors.New("unsupported map type")
	}
	return convertedMap, nil
}
