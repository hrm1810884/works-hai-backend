package usecase

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/domain/repository"
)

type RebootUserUsecase struct {
	userRepository    repository.UserRepository
	drawingRepository repository.DrawingRepository
}

func NewRebootUserUsecase(userRepository repository.UserRepository, drawingRepository repository.DrawingRepository) (*RebootUserUsecase, error) {
	return &RebootUserUsecase{userRepository, drawingRepository}, nil
}

func (u *RebootUserUsecase) RebootUser() error {
	id := uuid.New()
	startAiId, err := user.NewUserId(id)
	if err != nil {
		return err
	}
	startAiPosition := user.NewPosition(0, 0)

	staticAiImage, err := os.ReadFile("static/startAi.png")
	if err != nil {
		return err
	}

	startAiUrl, err := u.drawingRepository.UploadDrawing(startAiId.GetDrawingName(), staticAiImage)
	if err != nil {
		return err
	}

	startAi := user.NewUser(*startAiId, *startAiPosition, startAiUrl, true, time.Now(), time.Now())
	err = u.userRepository.Create(*startAi)
	if err != nil {
		return err
	}

	return nil
}
