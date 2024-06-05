package usecase

import (
	"context"
	"fmt"

	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/repository"
)

type IFetchPresignedUrls interface {
	FetchPresignedUrl() (map[string]string, error)
}

type FetchPresignedUrlsUsecase struct {
	storageClient *repository.FirebaseStorage
}

func NewFetchPresignedUrlsUsecase(ctx context.Context) (IFetchPresignedUrls, error) {
	firebaseApp, err := config.InitializeApp()
	if err != nil {
		panic(err) //FIXME: err handling 統一
	}

	storageClient, err := repository.NewFirebaseStorage(ctx, firebaseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage client: %w", err)
	}

	return &FetchPresignedUrlsUsecase{
		storageClient: storageClient,
	}, nil
}

func (u *FetchPresignedUrlsUsecase) FetchPresignedUrl() (map[string]string, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	bucketName := cfg.Firebase.Bucket
	presignedUrl, err := u.storageClient.GenerateSignedURL(bucketName, "hoge", 15)
	if err != nil {
		return nil, fmt.Errorf("error generating url: %w", err)
	}

	fetchedPresignedUrls := make(map[string]string)
	fetchedPresignedUrls["humanDrawing"] = presignedUrl

	return fetchedPresignedUrls, nil
}
