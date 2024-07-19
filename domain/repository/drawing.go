package repository

type DrawingRepository interface {
	GenerateSignedUrl() error
	DownloadImage() error
}
