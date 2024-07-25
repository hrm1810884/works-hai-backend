package usecase

import (
	"github.com/hrm1810884/works-hai-backend/domain"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/domain/repository"
)

type GetViewDataUsecase struct {
	repository repository.UserRepository
}

func NewGetViewDataUsecase(repository repository.UserRepository) (*GetViewDataUsecase, error) {
	return &GetViewDataUsecase{repository}, nil
}

func (u *GetViewDataUsecase) GetViewData(posX, posY int) (string, error) {
	targetPosition := user.NewPosition(posX, posY)
	user, err := u.repository.FindByPos(*targetPosition)
	if err != nil {
		return "", err
	}

	if !user.IsDrawn() {
		return "", domain.ErrNoLatestUser
	}

	url := user.GetUrl()
	return url, nil
}
