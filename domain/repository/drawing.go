package repository

type DrawingRepository interface {
	GenerateSignedUrl(drawingName string, method string) (string, error)
	DownloadDrawing(url string) (data []byte, err error)
	UploadDrawing(fileName string, fileData []byte) (string, error)
}
