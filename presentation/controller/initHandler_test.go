package controller_test

import (
	"context"
	"os"
	"testing"

	"github.com/hrm1810884/works-hai-backend/ogen"
	"github.com/hrm1810884/works-hai-backend/presentation/controller"
	"github.com/stretchr/testify/assert"
)

func TestInitGet(t *testing.T) {
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

	h := &controller.HaiHandler{}

	res, err := h.InitGet(ctx)
	assert.NoError(t, err, "InitGet failed")
	assert.NotNil(t, res, "Response should not be nil")

	switch v := res.(type) {
	case *ogen.InitGetOK:
		assert.NotEmpty(t, v.Result, "Result should not be empty")
		assert.NotEmpty(t, v.Result.ID, "Id should not be empty")
		assert.NotEmpty(t, v.Result.Urls, "Urls should not be empty")
		assert.NotEmpty(t, v.Result.Urls.HumanDrawing, "human should not be empty")
		t.Logf("HumanDrawing: %v", v.Result.Urls.HumanDrawing)
		t.Logf("BottomDrawing: %v", v.Result.Urls.BottomDrawing)
		t.Logf("RightDrawing: %v", v.Result.Urls.RightDrawing)
		t.Logf("LeftDrawing: %v", v.Result.Urls.LeftDrawing)
	case *ogen.InitGetBadRequest:
		t.Errorf("expected success but got bad request: %s", v.Error.Value)
	default:
		t.Errorf("unexpected response type: %T", v)
	}
}
