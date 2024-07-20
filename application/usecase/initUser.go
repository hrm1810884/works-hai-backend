package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/application/usecase/service"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/domain/repository"
)

type InitUserUsecase struct {
	Repository repository.UserRepository
	Service    service.GetSignedUrlService
}

func NewInitUserUsercase(service service.GetSignedUrlService, repository repository.UserRepository) (*InitUserUsecase, error) {
	return &InitUserUsecase{Repository: repository, Service: service}, nil
}

func (u *InitUserUsecase) InitUser(posX int, posY int) (url string, id string, err error) {
	currentPosition := user.NewPosition(posX, posY)
	userId, err := user.NewUserId(uuid.New())
	if err != nil {
		return "", "", fmt.Errorf("failed to get userId in InitUser: %w", err)
	}
	drawingName := userId.GetDrawingName()

	urlForGet, err := u.Service.GetSignedUrl(drawingName, "get")
	if err != nil {
		return "", "", fmt.Errorf("failed to get url for get: %w", err)
	}

	urlForPost, err := u.Service.GetSignedUrl(drawingName, "post")
	if err != nil {
		return "", "", fmt.Errorf("failed to get url for post: %w", err)
	}

	now := time.Now()
	newUser := user.NewUser(*userId, *currentPosition, urlForGet, now, now)
	err = u.Repository.Create(*newUser)
	if err != nil {
		return "", "", fmt.Errorf("failed to init user in db: %w", err)
	}

	return urlForPost, userId.ToId(), nil
}
