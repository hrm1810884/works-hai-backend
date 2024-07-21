package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/application/usecase/service"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/domain/repository"
)

type GenerateDrawingUsecase struct {
	userRepository    repository.UserRepository
	drawingRepository repository.DrawingRepository
	urlService        service.GetSignedUrlService
	generateService   service.DrawingService
}

func NewGenerateDrawingUsecase(ur repository.UserRepository, dr repository.DrawingRepository, urlService service.GetSignedUrlService, generateService service.DrawingService) (*GenerateDrawingUsecase, error) {
	return &GenerateDrawingUsecase{
		userRepository: ur, drawingRepository: dr, urlService: urlService, generateService: generateService,
	}, nil
}

func (u *GenerateDrawingUsecase) GenerateAIDrawing(userId user.UserId) (drawingUrl string, err error) {
	userData, err := u.userRepository.FindById(userId)
	if err != nil {
		return "", fmt.Errorf("not found user by id: %w", err)
	}

	aiPosition, err := userData.GetPosition().GetNext()
	if err != nil {
		return "", fmt.Errorf("failed to get next ai position: %w", err)
	}

	generatedDrawing, err := u.generateService.GenerateDrawing(aiPosition)
	if err != nil {
		return "", fmt.Errorf("failed to generate drawing: %w", err)
	}

	aiId, err := user.NewUserId(uuid.New())
	if err != nil {
		return "", fmt.Errorf("failed to get userId for ai generation: %w", err)
	}

	drawingUrl, err = u.drawingRepository.UploadDrawing(aiId.GetDrawingName(), generatedDrawing)
	if err != nil {
		return "", fmt.Errorf("failed to upload ai drawing:%w", err)
	}

	now := time.Now()
	aiData := user.NewUser(*aiId, *aiPosition, drawingUrl, now, now)
	err = u.userRepository.Create(*aiData)
	if err != nil {
		return "", fmt.Errorf("failed to create ai data in db: %w", err)
	}

	return drawingUrl, nil
}
