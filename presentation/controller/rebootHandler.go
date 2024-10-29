package controller

import (
	"context"

	"github.com/hrm1810884/works-hai-backend/application/usecase"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (h *HaiHandler) RebootPost(ctx context.Context) (ogen.RebootPostRes, error) {
	historyRepository, err := impl_repository.NewImplHistoryRepository(ctx)
	if err != nil {
		return &ogen.RebootPostBadRequest{Error: ogen.NewOptString("failed to get history repository")}, err
	}

	rebootHistoryUsecase := usecase.NewRebootHistoryUsecase(historyRepository)
	newHistory, err := rebootHistoryUsecase.RebootHistory()
	if err != nil {
		return &ogen.RebootPostBadRequest{Error: ogen.NewOptString("failed to reboot")}, err
	}

	userRepository, err := impl_repository.NewImplUserRepository(ctx, newHistory.GetHistoryId())
	if err != nil {
		return &ogen.RebootPostBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}
	drawingRepository, err := impl_repository.NewImplDrawingRepository(ctx)
	if err != nil {
		return &ogen.RebootPostBadRequest{Error: ogen.NewOptString("failed to get drawing repository")}, err
	}

	rebootUserUsecase, err := usecase.NewRebootUserUsecase(userRepository, drawingRepository)
	if err != nil {
		return &ogen.RebootPostBadRequest{Error: ogen.NewOptString("failed to get reboot usecase")}, err
	}

	err = rebootUserUsecase.RebootUser()
	if err != nil {
		return &ogen.RebootPostBadRequest{Error: ogen.NewOptString("failed to create init data")}, err
	}

	return &ogen.RebootPostOK{}, nil
}
