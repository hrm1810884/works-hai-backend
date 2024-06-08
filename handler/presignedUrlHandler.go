package handler

import (
	"context"
	"fmt"

	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/usecase"
)

func (h *HaiHandler) PresignedUrlsGet(ctx context.Context) (ogen.PresignedUrlsGetRes, error) {
	fetchedPresignedUrlsUsecase, err := usecase.NewFetchPresignedUrlsUsecase(ctx)
	if err != nil {
		return &ogen.PresignedUrlsGetBadRequest{}, fmt.Errorf("create presigned url usecase error: %w", err)
	}

	fetchedUrls, err := fetchedPresignedUrlsUsecase.FetchPresignedUrl()
	if err != nil {
		return &ogen.PresignedUrlsGetBadRequest{}, fmt.Errorf("fetch presigned url usecase error: %w", err)
	}

	return &ogen.PresignedUrlsGetOK{
		Result: ogen.PresignedUrlsGetOKResult{
			HumanDrawing: fetchedUrls["humanDrawing"],
		},
	}, nil
}
