package auth

import (
	"context"
	"errors"

	"github.com/hrm1810884/works-hai-backend/ogen"
)

type HaiSecurityHandler struct{}

func (m *HaiSecurityHandler) HandleApiKeyAuth(ctx context.Context, operationName string, auth ogen.ApiKeyAuth) (context.Context, error) {
	// 認証ロジックをここに実装
	if auth.APIKey == "hogehoge" {
		return ctx, nil
	}
	return nil, errors.New("invalid API key")
}
