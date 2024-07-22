package impl_repository

import (
	"context"
	"fmt"

	fs "firebase.google.com/go/v4/storage"
	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/infrastructure/impl/scripts"
	"github.com/hrm1810884/works-hai-backend/infrastructure/impl/storage"
)

type ImplDrawingRepository struct {
	Client *fs.Client
}

func NewImplDrawingRepository(ctx context.Context) (*ImplDrawingRepository, error) {
	app, err := config.InitializeApp()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase Storage client: %w", err)
	}

	return &ImplDrawingRepository{Client: client}, nil
}

func (dr *ImplDrawingRepository) GenerateSignedUrl(fileName string, method string) (url string, err error) {
	url, err = storage.ImplGenerateSignedUrl(dr.Client, fileName, method)
	if err != nil {
		return "", err
	}
	return url, nil
}
func (dr *ImplDrawingRepository) UploadDrawing(fileName string, fileData []byte) (url string, err error) {
	err = storage.ImplUploadDrawing(dr.Client, fileName, fileData)
	if err != nil {
		return "", err
	}

	// 署名付きURLの生成
	url, err = dr.GenerateSignedUrl(fileName, "GET")
	if err != nil {
		return "", err
	}

	return url, nil

}

func (dr *ImplDrawingRepository) DownloadDrawing(url string) (data []byte, err error) {
	data, err = storage.ImplDownloadDrawing(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dr *ImplDrawingRepository) GenerateAIDrawing(drawingData map[string][]byte) (data []byte, err error) {
	data, err = scripts.ImplGenerateAIDrawing(drawingData)
	if err != nil {
		return nil, err
	}
	return data, nil
}
