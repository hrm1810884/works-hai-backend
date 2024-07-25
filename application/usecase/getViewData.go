package usecase

import (
	"fmt"

	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/domain/repository"
)

type GetViewDataUsecase struct {
	repository repository.UserRepository
}

func NewGetViewDataUsecase(repository repository.UserRepository) (*GetViewDataUsecase, error) {
	return &GetViewDataUsecase{repository}, nil
}

func (u *GetViewDataUsecase) GetViewData() ([]user.User, error) {
	arr, err := u.repository.GetLatestArray()
	if err != nil {
		return nil, err
	}

	if len(arr) == 0 {
		return nil, fmt.Errorf("no data as latest array")
	}

	return arr, nil
}
