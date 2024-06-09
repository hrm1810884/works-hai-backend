package handler

import (
	"context"

	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (h *HaiHandler) PresignedUrlsGet(ctx context.Context) (ogen.PresignedUrlsGetRes, error) {
	// fetchedPresignedUrlsservice, err := service.NewGetSignedUrlService(ctx)
	// if err != nil {
	// 	return &ogen.PresignedUrlsGetBadRequest{}, fmt.Errorf("create presigned url service error: %w", err)
	// }

	// fetchedUrls, err := fetchedPresignedUrlsservice.GetSignedUrl("PUT") //FIXME: reqにposition情報かな
	// if err != nil {
	// 	return &ogen.PresignedUrlsGetBadRequest{}, fmt.Errorf("fetch presigned url service error: %w", err)
	// }

	// return &ogen.PresignedUrlsGetOK{
	// 	Result: ogen.PresignedUrlsGetOKResult{
	// 		HumanDrawing: fetchedUrls["humanDrawing"],
	// 	},
	// }, nil

	return &ogen.PresignedUrlsGetOK{}, nil
}
