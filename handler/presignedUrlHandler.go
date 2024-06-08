package handler

import (
	"context"
	"fmt"

	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/usecase/service"
)

func (h *HaiHandler) PresignedUrlsGet(ctx context.Context) (ogen.PresignedUrlsGetRes, error) {
	fetchedPresignedUrlsservice, err := service.NewFetchPresignedUrlsService(ctx)
	if err != nil {
		return &ogen.PresignedUrlsGetBadRequest{}, fmt.Errorf("create presigned url service error: %w", err)
	}

	fetchedUrls, err := fetchedPresignedUrlsservice.FetchPresignedUrl("PUT")
	if err != nil {
		return &ogen.PresignedUrlsGetBadRequest{}, fmt.Errorf("fetch presigned url service error: %w", err)
	}

	return &ogen.PresignedUrlsGetOK{
		Result: ogen.PresignedUrlsGetOKResult{
			HumanDrawing: fetchedUrls["humanDrawing"],
		},
	}, nil
}
