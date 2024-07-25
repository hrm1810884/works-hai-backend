package controller

import (
	"context"

	"github.com/hrm1810884/works-hai-backend/application/usecase"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (h *HaiHandler) ViewGet(ctx context.Context) (ogen.ViewGetRes, error) {
	userRepository, err := impl_repository.NewImplUserRepository(ctx)
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	viewerUsecase, err := usecase.NewGetViewDataUsecase(userRepository)
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get usecase")}, err
	}

	arr, err := viewerUsecase.GetViewData()
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get view data")}, err
	}

	var resArr []ogen.ViewGetOKResultItem
	for i := 0; i < len(arr); i++ {
		resItem := &ogen.ViewGetOKResultItem{
			Position: ogen.ViewGetOKResultItemPosition{
				X: arr[i].GetPosition().GetX(),
				Y: arr[i].GetPosition().GetY(),
			},
			URL: arr[i].GetUrl(),
		}
		resArr = append(resArr, *resItem)
	}

	return &ogen.ViewGetOK{Result: resArr}, nil
}
