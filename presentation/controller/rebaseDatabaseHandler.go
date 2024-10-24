package controller

import (
	context "context"

	usecase "github.com/hrm1810884/works-hai-backend/application/usecase"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
	ogen "github.com/hrm1810884/works-hai-backend/ogen"
	"golang.org/x/text/number"
)

func (h *HaiHandler) RebaseDatabaseGet(ctx context.Context) (ogen.RebaseDatabaseGetRes, error) {
	// initialize impl_userRepository
	var implUserRepository *impl_repository.ImplUserRepository
	var implDrawingRepository * impl_repository.ImplDrawingRepository
	var err error
	
	implUserRepository, err = impl_repository.NewImplUserRepository(ctx)
	if err != nil {
		return &ogen.RebaseDatabaseGetBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	rebaseDatabaseUsecase, err := usecase.NewRebaseDatabaseUsecase(implUserRepository, implDrawingRepository)
	if err != nil {
		return &ogen.RebaseDatabaseGetBadRequest{Error: ogen.NewOptString("failed to get init usecase")}, err
	}

	err = rebaseDatabaseUsecase.DeleteImagesExceptCenter()
	if err != nil {
		return &ogen.RebaseDatabaseGetOK{
			Success: false,
		}, nil
	}

	err = rebaseDatabaseUsecase.OverrideInitialImage()
	if err != nil {
		return &ogen.RebaseDatabaseGetOK{
			Success: false,
		}, nil
	}

	return &ogen.RebaseDatabaseGetOK{
		Success: true,
	}, nil
}
