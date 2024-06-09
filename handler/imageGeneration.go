package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/hrm1810884/works-hai-backend/entity"
	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/usecase"
)

func (*HaiHandler) ImageGenerationPost(ctx context.Context, req *ogen.ImageGenerationPostReq) (ogen.ImageGenerationPostRes, error) {
	log.Print("hoge")
	reqPosition := entity.NewPositionEntity(req.X.Value, req.Y.Value)
	getSignedUrlsUsecase, err := usecase.NewGetSignedUrlsUsecase(reqPosition)
	if err != nil {
		return nil, fmt.Errorf("usecase error: get presigned urls")
	}
	message, err := getSignedUrlsUsecase.GenerateAIDrawing(ctx)
	if err != nil {
		return &ogen.ImageGenerationPostBadRequest{Error: ogen.NewOptString("failed to generate image...")}, err
	}

	log.Printf("image generation succeeded: %q", message)

	return &ogen.ImageGenerationPostOK{Message: ogen.NewOptString(message)}, nil
}
