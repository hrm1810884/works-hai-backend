package controller_test

import (
	"context"
	"os"
	"testing"

	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/presentation/controller"
	"github.com/stretchr/testify/assert"
)

func TestViewHandler(t *testing.T) {
	t.Parallel()
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
	req := &ogen.ViewGetReq{
		Position: ogen.ViewGetReqPosition{
			X: 0,
			Y: 0,
		},
	}

	// コントローラーのハンドラを作成
	h := &controller.HaiHandler{}

	// 実行
	res, err := h.ViewGet(ctx, req)
	assert.NoError(t, err, "ImageGenerationPost failed")
	assert.NotNil(t, res, "Response should not be nil")

	// レスポンスの検証
	switch v := res.(type) {
	case *ogen.ViewGetOK:
		println("%s", v.Result.URL)
		assert.NotEmpty(t, v.Result.URL, "URL should not be empty")
	case *ogen.ViewGetBadRequest:
		t.Errorf("expected success but got bad request: %s", v.Error.Value)
	default:
		t.Errorf("unexpected response type: %T", v)
	}
}
