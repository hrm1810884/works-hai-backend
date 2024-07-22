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

func ImplGenerateSignedUrl(client *storage.Client, fileName string, method string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", fmt.Errorf("error loading config: %w", err)
	}

	bucketName := cfg.Firebase.Bucket

	// 署名付きURLのオプションを設定
	opts := &cs.SignedURLOptions{
		Scheme:  cs.SigningSchemeV4,
		Method:  method,
		Expires: time.Now().Add(7 * 24 * time.Hour), // 有効期限
	}

	// 署名付きURLを生成
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	url, err := bucket.SignedURL(fileName, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucketName, err)
	}

	return url, nil
}

func ImplUploadDrawing(client *storage.Client, fileName string, fileData []byte) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}
	bucketName := cfg.Firebase.Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("failed to get bucket: %w", err)
	}

	// ファイルのContentTypeを推測
	contentType := http.DetectContentType(fileData)

	// ファイルへの書き込み用のWriterを作成
	wc := bucket.Object(fileName).NewWriter(context.Background())
	wc.ContentType = contentType
	wc.CacheControl = "public, max-age=31536000" // 1年間キャッシュする

	if _, err := wc.Write(fileData); err != nil {
		return fmt.Errorf("failed to write image to Firebase Storage: %w", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	return nil
}

func ImplDownloadDrawing(url string) (data []byte, err error) {
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
