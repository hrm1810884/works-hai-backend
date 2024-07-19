package storage

import (
	"context"
	"fmt"
	"time"

	cs "cloud.google.com/go/storage"
	"firebase.google.com/go/v4/storage"
	"github.com/hrm1810884/works-hai-backend/config"
)

type ImplDrawingRepository struct {
	Client *storage.Client
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

func (dr *ImplDrawingRepository) GenerateSignedUrl(drawingName string, expiry time.Duration, method string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", fmt.Errorf("error loading config: %w", err)
	}

	bucketName := cfg.Firebase.Bucket

	// 署名付きURLのオプションを設定
	opts := &cs.SignedURLOptions{
		Scheme:  cs.SigningSchemeV4,
		Method:  method,
		Expires: time.Now().Add(expiry * time.Minute), // 有効期限
	}

	// 署名付きURLを生成
	bucket, err := dr.Client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	url, err := bucket.SignedURL(drawingName, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucketName, err)
	}

	return url, nil
}
