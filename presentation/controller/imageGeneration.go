package controller

import (
	"context"
	"log"

	"github.com/hrm1810884/works-hai-backend/application/usecase"
	"github.com/hrm1810884/works-hai-backend/domain/entity"
	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (*HaiHandler) ImageGenerationPost(ctx context.Context, req *ogen.ImageGenerationPostReq) (ogen.ImageGenerationPostRes, error) {
	log.Print("hoge")
	reqPosition := entity.NewPositionEntity(req.X.Value, req.Y.Value)
	currentPosition, err := reqPosition.GetNext()
	if err != nil {
		return &ogen.ImageGenerationPostBadRequest{Error: ogen.NewOptString("failed to go next")}, err
	}
	getSignedUrlsUsecase := usecase.NewGenerateImageUsecase(ctx, currentPosition)

	message, err := getSignedUrlsUsecase.GenerateAIDrawing(ctx)
	if err != nil {
		return &ogen.ImageGenerationPostBadRequest{Error: ogen.NewOptString("failed to generate image...")}, err
	}

	log.Printf("image generation succeeded: %q", message)

	return &ogen.ImageGenerationPostOK{Message: ogen.NewOptString(message)}, nil
}
