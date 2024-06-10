package handler

import (
	"context"

	"github.com/hrm1810884/works-hai-backend/entity"
	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/usecase"
)

func (h *HaiHandler) PresignedUrlsGet(ctx context.Context) (ogen.PresignedUrlsGetRes, error) {
	reqPosition := entity.NewPositionEntity(0, -1)
	currentPosition, err := reqPosition.GetNext()
	if err != nil {
		return &ogen.PresignedUrlsGetBadRequest{}, err
	}

	u := usecase.NewQuadSignedUrlsUsecase(ctx, currentPosition)

	res, err := u.GenerateQuadSignedUrls(ctx, currentPosition)
	if err != nil {
		return &ogen.PresignedUrlsGetBadRequest{}, err
	}

	return &ogen.PresignedUrlsGetOK{
		Result: ogen.PresignedUrlsGetOKResult{
			HumanDrawing:  res["Human"],
			TopDrawing:    ogen.NewOptString(res["Top"]),
			RightDrawing:  ogen.NewOptString(res["Right"]),
			BottomDrawing: ogen.NewOptString(res["Bottom"]),
			LeftDrawing:   ogen.NewOptString(res["Left"]),
		},
	}, nil
}
