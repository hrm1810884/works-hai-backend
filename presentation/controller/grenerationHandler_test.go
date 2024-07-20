package controller_test

import (
	"context"
	"os"
	"testing"

	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/presentation/controller"
	"github.com/stretchr/testify/assert"
)

func TestImageGenerationPost(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	t.Cleanup(func() {
		err := os.Chdir(cwd)
		if err != nil {
			t.Fatalf("failed to get current working directory: %v", err)
		}
	})
	err = os.Chdir("../..")
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	ctx := context.Background()

	// テスト用のリクエストデータを作成
	req := &ogen.GeneratePostReq{
		UserId: "78479344-30ab-4f5f-8f1c-f10ba5064598",
	}

	// コントローラーのハンドラを作成
	h := &controller.HaiHandler{}

	// 実行
	res, err := h.ImageGenerationPost(ctx, req)
	assert.NoError(t, err, "ImageGenerationPost failed")
	assert.NotNil(t, res, "Response should not be nil")

	// レスポンスの検証
	switch v := res.(type) {
	case *ogen.GeneratePostOK:
		assert.NotEmpty(t, v.URL, "URL should not be empty")
	case *ogen.GeneratePostBadRequest:
		t.Errorf("expected success but got bad request: %s", v.Error.Value)
	default:
		t.Errorf("unexpected response type: %T", v)
	}
}
