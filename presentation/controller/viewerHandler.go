package controller

import (
	"context"
	"errors"

	"github.com/hrm1810884/works-hai-backend/application/usecase"
	"github.com/hrm1810884/works-hai-backend/domain"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (h *HaiHandler) ViewGet(ctx context.Context, req *ogen.ViewGetReq) (ogen.ViewGetRes, error) {
	posX := req.Position.X
	posY := req.Position.Y

	userRepository, err := impl_repository.NewImplUserRepository(ctx)
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	viewerUsecase, err := usecase.NewGetViewDataUsecase(userRepository)
	if err != nil {
		return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get usecase")}, err
	}

	url, err := viewerUsecase.GetViewData(posX, posY)
	switch {
	case errors.Is(err, domain.ErrNoLatestUser):
		{
			return &ogen.ViewGetNotFound{Error: ogen.NewOptString("no need to update")}, err
		}
	case err != nil:
		{
			return &ogen.ViewGetBadRequest{Error: ogen.NewOptString("failed to get viewer data")}, err
		}
	default:
		{
			return &ogen.ViewGetOK{Result: ogen.ViewGetOKResult{Position: ogen.ViewGetOKResultPosition{X: posX, Y: posY}, URL: url}}, nil
		}
	}

}
