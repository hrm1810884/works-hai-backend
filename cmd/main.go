package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hrm1810884/works-hai-backend/ogen" // パッケージをインポート
)

// MyUploadService は UploadURLGetRes インターフェースを実装するサービスです。
type MyUploadService struct{}

func (s *MyUploadService) UploadURLGet(ctx context.Context) (ogen.UploadURLGetRes, error) {
	// URL 生成ロジックをここに記述
	// 例として固定のURLを返す
	select {
	case <-time.After(1 * time.Second): // シミュレーションのための遅延
		uploadURL := "https://example-cloud-storage.com/user_drawing.png?signature=..."
		return &ogen.UploadURLGetOK{PresignedURL: ogen.OptString{Value: uploadURL, Set: true}}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	uploadService := &MyUploadService{}

	http.HandleFunc("/upload-url", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		res, err := uploadService.UploadURLGet(ctx)
		if err != nil {
			http.Error(w, "Invalid input data", http.StatusBadRequest)
			return
		}

		switch res := res.(type) {
		case *ogen.UploadURLGetOK:
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(res); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
		case *ogen.UploadURLGetBadRequest:
			http.Error(w, "Invalid input data", http.StatusBadRequest)
		default:
			http.Error(w, "Unknown error", http.StatusInternalServerError)
		}
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Printf("Server is running on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
