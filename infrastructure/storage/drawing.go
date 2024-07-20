package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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

func (dr *ImplDrawingRepository) UploadDrawing(fileName string, fileData []byte) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", fmt.Errorf("error loading config: %w", err)
	}
	bucketName := cfg.Firebase.Bucket
	bucket, err := dr.Client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("failed to get bucket: %v", err)
	}

	// ファイルのContentTypeを推測
	contentType := http.DetectContentType(fileData)

	// ファイルへの書き込み用のWriterを作成
	wc := bucket.Object(fileName).NewWriter(context.Background())
	wc.ContentType = contentType
	wc.CacheControl = "public, max-age=31536000" // 1年間キャッシュする

	if _, err := wc.Write(fileData); err != nil {
		return "", fmt.Errorf("failed to write image to Firebase Storage: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// 署名付きURLの生成
	signedURL, err := dr.GenerateSignedUrl(fileName, 15, "GET")
	if err != nil {
		return "", err
	}

	return signedURL, nil
}

func (dr *ImplDrawingRepository) DownloadDrawing(url, filePath string) (data []byte, err error) {
	validUrl, err := validateURL(url)
	if err != nil {
		// Handle the error appropriately
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	resp, err := http.Get(validUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Print("%w", url)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: received non-200 response code %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func validateURL(u string) (*url.URL, error) {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, fmt.Errorf("invalid URL scheme")
	}

	return parsedURL, nil
}
