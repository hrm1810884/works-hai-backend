package service

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/hrm1810884/works-hai-backend/ogen"
)

func (h *HaiHandler) AiDrawingGet(ctx context.Context) (ogen.AiDrawingGetRes, error) {
	fmt.Println("Received get ai drawing request")

	// Pythonスクリプトを実行するためのコマンドを作成
	cmd := exec.Command("python3", "hello.py")
	var out bytes.Buffer
	cmd.Stdout = &out

	// コマンドを非同期で実行する
	errChan := make(chan error)
	go func() {
		errChan <- cmd.Run()
	}()

	// Pythonスクリプトの実行を待つ間に他の処理を行う
	fmt.Println("Doing other processing while waiting for the script to finish...")
	time.Sleep(2 * time.Second) // 仮の処理として2秒間待機

	// コマンドの結果を待つ
	err := <-errChan
	if err != nil {
		return &ogen.AiDrawingGetBadRequest{}, fmt.Errorf("failed to execute python script: %w", err)
	}

	// Pythonの出力をレスポンスとして返す
	result := ogen.AiDrawingGetOK{
		Result: ogen.AiDrawingGetOKResult{
			TopDrawing: ogen.OptString{
				Value: out.String(),
				Set:   true,
			},
		},
	}
	return &result, nil
}
