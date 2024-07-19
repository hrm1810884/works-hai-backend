package repository

type DrawingRepository interface {
	GetSignedUrls() error
	DownloadImage() error
}
