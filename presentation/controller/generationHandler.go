package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/application/usecase"
	"github.com/hrm1810884/works-hai-backend/application/usecase/service"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (*HaiHandler) GeneratePost(ctx context.Context, req *ogen.GeneratePostReq) (ogen.GeneratePostRes, error) {
	reqId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("invalid request to convert to uuid")}, err
	}

	historyRepository, err := impl_repository.NewImplHistoryRepository(ctx)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get history repository")}, err
	}

	currentHistory, err := historyRepository.FindLatest()
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get current history")}, err
	}

	userId, err := user.NewUserId(reqId)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get userId")}, err
	}

	userRepository, err := impl_repository.NewImplUserRepository(ctx, currentHistory.GetHistoryId())
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	drawingRepository, err := impl_repository.NewImplDrawingRepository(ctx)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get drawing repository")}, err
	}

	generateService, err := service.NewDrawingService(userRepository, drawingRepository)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get generate service")}, err
	}

	urlService, err := service.NewGetSignedUrlService(drawingRepository)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get url service")}, err
	}

	generateUsecase, err := usecase.NewGenerateDrawingUsecase(userRepository, drawingRepository, *urlService, *generateService)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get usecase")}, err
	}

	url, err := generateUsecase.GenerateAIDrawing(*userId)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to generate image...")}, err
	}

	return &ogen.GeneratePostOK{URL: url}, nil
}
