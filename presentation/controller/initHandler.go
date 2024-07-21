package controller

import (
	"context"
	"errors"

	"github.com/hrm1810884/works-hai-backend/application/usecase"
	"github.com/hrm1810884/works-hai-backend/application/usecase/service"
	"github.com/hrm1810884/works-hai-backend/domain"
	"github.com/hrm1810884/works-hai-backend/infrastructure/database"
	"github.com/hrm1810884/works-hai-backend/infrastructure/storage"
	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (h *HaiHandler) InitGet(ctx context.Context) (ogen.InitGetRes, error) {
	userRepository, err := database.NewImplUserRepository(ctx)
	if err != nil {
		return &ogen.InitGetBadRequest{Error: ogen.NewOptString("failed to get user repository")}, err
	}

	drawingRepository, err := storage.NewImplDrawingRepository(ctx)
	if err != nil {
		return &ogen.InitGetBadRequest{Error: ogen.NewOptString("failed to get drawing repository")}, err
	}

	urlService, err := service.NewGetSignedUrlService(drawingRepository)
	if err != nil {
		return &ogen.InitGetBadRequest{Error: ogen.NewOptString("failed to get url service")}, err
	}

	drawingService, err := service.NewDrawingService(userRepository, drawingRepository)
	if err != nil {
		return &ogen.InitGetBadRequest{Error: ogen.NewOptString("failed to get url service")}, err
	}

	initUserUsecase, err := usecase.NewInitUserUsercase(userRepository, *urlService, *drawingService)
	if err != nil {
		return &ogen.InitGetBadRequest{Error: ogen.NewOptString("failed to get init usecase")}, err
	}

	var posX, posY int
	latestUser, err := userRepository.FindLatest()
	switch {
	case err == nil:
		pos := latestUser.GetPosition()
		posX = pos.GetX()
		posY = pos.GetY()
	case errors.Is(err, domain.ErrNoLatestUser):
		posX = 0
		posY = 0
	default:
		return &ogen.InitGetBadRequest{Error: ogen.NewOptString("failed to get latest user")}, err
	}

	urls, id, err := initUserUsecase.InitUser(posX, posY)
	if err != nil {
		return &ogen.InitGetBadRequest{Error: ogen.NewOptString("failed to init user")}, err
	}

	return &ogen.InitGetOK{
		Result: ogen.InitGetOKResult{
			ID: id,
			Urls: ogen.InitGetOKResultUrls{
				HumanDrawing:  urls["human"],
				TopDrawing:    convertUrlToOptString(urls["top"]),
				RightDrawing:  convertUrlToOptString(urls["right"]),
				BottomDrawing: convertUrlToOptString(urls["bottom"]),
				LeftDrawing:   convertUrlToOptString(urls["left"]),
			},
		},
	}, nil
}

func convertUrlToOptString(url string) ogen.OptString {
	opt := ogen.NewOptString(url)
	if url == "" {
		opt.Reset() // urlが空文字の場合はリセット
		return opt
	}
	return opt
}
