package repository

type DrawingRepository interface {
	GenerateSignedUrl(fileName string, method string) (url string, err error)
	GenerateAIDrawing(drawingData map[string][]byte) (data []byte, err error)
	UploadDrawing(fileName string, fileData []byte) (url string, err error)
	DownloadDrawing(url string) (data []byte, err error)
}
