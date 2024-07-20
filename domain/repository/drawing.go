package repository

import "time"

type DrawingRepository interface {
	GenerateSignedUrl(drawingName string, expiry time.Duration, method string) (string, error)
}
