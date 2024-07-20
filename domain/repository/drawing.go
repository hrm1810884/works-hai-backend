package repository

import (
	"time"
)

type DrawingRepository interface {
	GenerateSignedUrl(drawingName string, expiry time.Duration, method string) (string, error)
	DownloadDrawing(url string) (data []byte, err error)
	UploadDrawing(fileName string, fileData []byte) (string, error)
	GenerateAIDrawing(drawingData map[string][]byte) ([]byte, error)
}
