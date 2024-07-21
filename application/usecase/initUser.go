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
	repository     repository.UserRepository
	urlService     service.GetSignedUrlService
	drawingService service.DrawingService
}

func NewInitUserUsercase(repository repository.UserRepository, urlService service.GetSignedUrlService, drawingService service.DrawingService) (*InitUserUsecase, error) {
	return &InitUserUsecase{repository, urlService, drawingService}, nil
}

func (u *InitUserUsecase) InitUser(posX int, posY int) (urls map[string]string, id string, err error) {
	currentPosition := user.NewPosition(posX, posY)
	userId, err := user.NewUserId(uuid.New())
	if err != nil {
		return nil, "", fmt.Errorf("failed to get userId in InitUser: %w", err)
	}
	drawingName := userId.GetDrawingName()

	quadUrls, err := u.drawingService.GetQuadUrls(*currentPosition)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get quad urls for (%d, %d): %w", posX, posY, err)
	}

	urlForGet, err := u.urlService.GetSignedUrl(drawingName, "get")
	if err != nil {
		return nil, "", fmt.Errorf("failed to get url for get: %w", err)
	}

	urlForPost, err := u.urlService.GetSignedUrl(drawingName, "put")
	if err != nil {
		return nil, "", fmt.Errorf("failed to get url for post: %w", err)
	}

	now := time.Now()
	newUser := user.NewUser(*userId, *currentPosition, urlForGet, now, now)
	err = u.repository.Create(*newUser)
	if err != nil {
		return nil, "", fmt.Errorf("failed to init user in db: %w", err)
	}

	urls = map[string]string{
		"human":  urlForPost,
		"top":    quadUrls["top"],
		"right":  quadUrls["right"],
		"bottom": quadUrls["bottom"],
		"left":   quadUrls["left"],
	}

	return urls, userId.ToId(), nil
}
