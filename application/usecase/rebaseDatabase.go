package usecase

import (
	"time"
	"fmt"
	"github.com/google/uuid"
	user "github.com/hrm1810884/works-hai-backend/domain/entity/user"
	repository "github.com/hrm1810884/works-hai-backend/domain/repository"
)

type RebaseDatabaseUsecase struct {
	userRepository    repository.UserRepository
	drawingRepository repository.DrawingRepository
}

func NewRebaseDatabaseUsecase( ur repository.UserRepository ,dr repository.DrawingRepository) (
	*RebaseDatabaseUsecase, error) {
	return &RebaseDatabaseUsecase{ userRepository: ur , drawingRepository: dr}, nil
}

func (usecase *RebaseDatabaseUsecase) DeleteImagesExceptCenter() (err error) {
	err = usecase.userRepository.DeleteImagesExceptCenter();
	if err != nil {
		return err
	}

	return nil
}

func (usecase *RebaseDatabaseUsecase) OverrideInitialImage() (err error) {
	// generate image from scratch
	emptyData := map[string][]byte{}
	data, err := usecase.drawingRepository.GenerateAIDrawing(emptyData)
	if err != nil {
		return fmt.Errorf("failed to generate drawing: %w", err)
	}

	// override intial image with new generated image
	centerPosition := user.NewPosition(0, 0)
	aiData, err := usecase.userRepository.FindByPos(*centerPosition)
	if err != nil {
		newAiId, err := user.NewUserId(uuid.New())
		if err != nil {
			return fmt.Errorf("failed to instantiate userId: %w", err)
		}

		drawingUrl, err := usecase.drawingRepository.UploadDrawing(newAiId.GetDrawingName(), data)
		if err != nil {
			return fmt.Errorf("failed to upload ai drawing:%w", err)
		}

		now := time.Now()
		aiDataToOverride := user.NewUser(*newAiId, *centerPosition, drawingUrl, true, now, now)
		err = usecase.userRepository.Create(*aiDataToOverride)
		if err != nil {
			return fmt.Errorf("failed to create ai data in db: %w", err)
		}

	}
	
	aiId := aiData.GetId()
	drawingUrl, err := usecase.drawingRepository.UploadDrawing(aiId.GetDrawingName(), data)
	if err != nil {
		return fmt.Errorf("failed to upload ai drawing:%w", err)
	}

	now := time.Now()
	aiDataToOverride := user.NewUser(*aiId, *centerPosition, drawingUrl, true, now, now)
	err = usecase.userRepository.Create(*aiDataToOverride)
	if err != nil {
		return fmt.Errorf("failed to create ai data in db: %w", err)
	}

	return nil
}
