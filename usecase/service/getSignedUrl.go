package service

import (
	"context"
	"fmt"

	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/repository"
)

type IGetSignedUrl interface {
	GetSignedUrl(resourceName string, method string) (string, error)
}

type GetSignedUrlService struct {
	storageClient *repository.FirebaseStorageRepository
}

func NewGetSignedUrlService(ctx context.Context) (IGetSignedUrl, error) {
	firebaseApp, err := config.InitializeApp()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	storageClient, err := repository.NewFirebaseStorageRepository(ctx, firebaseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage client: %w", err)
	}

	return &GetSignedUrlService{
		storageClient: storageClient,
	}, nil
}

func (s *GetSignedUrlService) GetSignedUrl(resourceName string, method string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", fmt.Errorf("error loading config: %w", err)
	}

	bucketName := cfg.Firebase.Bucket

	presignedUrl, err := s.storageClient.GenerateSignedURL(bucketName, resourceName, 15, method)
	if err != nil {
		return "", fmt.Errorf("error generating url: %w", err)
	}

	return presignedUrl, nil
}
