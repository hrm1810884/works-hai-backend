package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func uploadHandler(c echo.Context) error {
	log.Println("Received upload request")

	// 最大ファイルサイズを10MBに制限
	err := c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error parsing multipart form: %v", err))
	}

	file, handler, err := c.Request().FormFile("image")
	if err != nil {
		log.Printf("Error retrieving the file: %v", err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error retrieving the file: %v", err))
	}
	defer file.Close()

	log.Printf("Uploaded File: %s, Size: %d, MIME: %s", handler.Filename, handler.Size, handler.Header.Get("Content-Type"))

	// アップロードされたファイルをメモリに読み込む
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		log.Printf("Error reading the file: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading the file: %v", err))
	}

	// バイトデータを画像としてデコード
	img, _, err := image.Decode(&buf)
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error decoding image: %v", err))
	}

	// 保存するPNGファイルのパスを指定
	savePath := filepath.Join("uploads", handler.Filename+".png")

	// 保存するディレクトリが存在しない場合は作成
	if err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm); err != nil {
		log.Printf("Error creating directory: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating directory: %v", err))
	}

	// ファイルを保存
	destFile, err := os.Create(savePath)
	if err != nil {
		log.Printf("Error creating the file: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating the file: %v", err))
	}
	defer destFile.Close()

	// 画像をPNG形式でエンコードして保存
	err = png.Encode(destFile, img)
	if err != nil {
		log.Printf("Error encoding image to PNG: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error encoding image to PNG: %v", err))
	}

	log.Printf("File uploaded and converted successfully: %s", handler.Filename)
	return c.String(http.StatusCreated, fmt.Sprintf("File uploaded and converted successfully: %s", handler.Filename))
}

func main() {
	e := echo.New()

	// CORSミドルウェアを適用
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "x-api-key",
		},
	}))

	e.POST("/human-drawing", uploadHandler)

	log.Println("Server started at http://localhost:8080")
	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
