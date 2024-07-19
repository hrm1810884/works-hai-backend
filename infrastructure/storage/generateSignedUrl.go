package storage

import (
	"context"
	"fmt"
	"time"

	cs "cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
)

type FirebaseStorageRepository struct {
	Client *storage.Client
}

func NewFirebaseStorageRepository(ctx context.Context, app *firebase.App) (*FirebaseStorageRepository, error) {
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase Storage client: %w", err)
	}

	return &FirebaseStorageRepository{Client: client}, nil
}

func (tfs *FirebaseStorageRepository) GenerateSignedURL(bucketName, objectName string, expiry time.Duration, method string) (string, error) {
	// 署名付きURLのオプションを設定
	opts := &cs.SignedURLOptions{
		Scheme:  cs.SigningSchemeV4,
		Method:  method,
		Expires: time.Now().Add(expiry * time.Minute), // 有効期限
	}

	// 署名付きURLを生成
	bucket, err := tfs.Client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	u, err := bucket.SignedURL(objectName, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucketName, err)
	}

	return u, nil
}
