package domain

import "errors"

var (
	ErrNoLatestUser = errors.New("no latest user found")
	// 他のユーザー関連のエラーもここに定義できます
)
