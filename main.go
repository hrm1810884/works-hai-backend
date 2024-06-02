package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hrm1810884/works-hai-backend/auth"
	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/service"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ここでCORSを設定
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-api-key")

		// オプションズリクエストに対するプリフライト応答
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 次のハンドラーを呼び出す
		next.ServeHTTP(w, r)
	})
}

func main() {
	// サーバーの初期設定
	hdl, err := ogen.NewServer(
		&service.HaiHandler{},
		&auth.HaiSecurityHandler{},
	)

	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	// 設定の読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Printf("Server started at %v\n", cfg.Server.Dev)

	// サーバーの設定
	srv := &http.Server{
		Addr:        cfg.Server.Dev,
		Handler:     enableCORS(hdl),
		ReadTimeout: 30 * time.Second,
	}

	// サーバーの起動
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
