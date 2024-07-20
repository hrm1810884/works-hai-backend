package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/application/usecase"
	"github.com/hrm1810884/works-hai-backend/application/usecase/service"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/infrastructure/database"
	"github.com/hrm1810884/works-hai-backend/infrastructure/storage"
	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (*HaiHandler) ImageGenerationPost(ctx context.Context, req *ogen.GeneratePostReq) (ogen.GeneratePostRes, error) {
	reqId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("invalid request to convert to uuid")}, err
	}

	userId, err := user.NewUserId(reqId)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get userId")}, err
	}

	userRepository, err := database.NewImplUserRepository(ctx)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	drawingRepository, err := storage.NewImplDrawingRepository(ctx)
	if err != nil {
		return &ogen.GeneratePostBadRequest{Error: ogen.NewOptString("failed to get drawing repository")}, err
	}

	generateService, err := service.NewGenerateDrawingService(userRepository, drawingRepository)
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
