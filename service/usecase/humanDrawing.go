package usecase

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func SaveImage(file io.Reader) error {
	// アップロードされたファイルをメモリに読み込む
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return err
	}

	// バイトデータを画像としてデコード
	img, _, err := image.Decode(&buf)
	if err != nil {
		return err
	}

	// 保存するPNGファイルのパスを指定
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102_150405")
	filename := formattedTime + ".png"
	savePath := filepath.Join("uploads", filename)

	// 保存するディレクトリが存在しない場合は作成
	if err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm); err != nil {
		return err
	}

	// ファイルを保存
	destFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 画像をPNG形式でエンコードして保存
	err = png.Encode(destFile, img)
	if err != nil {
		return err
	}

	log.Printf("File uploaded and converted successfully: %s", filename)
	return nil
}
