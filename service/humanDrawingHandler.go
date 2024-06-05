package service

import (
	"context"
	"fmt"
	"log"

	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/service/usecase"
)

func (h *HaiHandler) HumanDrawingPost(ctx context.Context, req *ogen.HumanDrawingPostReq) (ogen.HumanDrawingPostRes, error) {
	fmt.Println("Received upload request")

	humandrawing := req.Image

	err := usecase.SaveImage(humandrawing.File)
	if err != nil {
		log.Printf("Error processing the image: %v", err)
		return &ogen.HumanDrawingPostBadRequest{}, fmt.Errorf("error processing the image: %w", err)
	}

	return &ogen.HumanDrawingPostOK{
		Message: ogen.NewOptString("File uploaded and converted successfully"),
	}, nil
}

func (h *HaiHandler) SavedURLPost(ctx context.Context, req *ogen.SavedURLPostReq) (ogen.SavedURLPostRes, error) {
	// 描画URLを保存するロジック
	return &ogen.SavedURLPostOK{}, nil
}

func (h *HaiHandler) UploadURLGet(ctx context.Context) (ogen.UploadURLGetRes, error) {
	// アップロード用のプリサインURLを取得するロジック
	return &ogen.UploadURLGetOK{}, nil
}
