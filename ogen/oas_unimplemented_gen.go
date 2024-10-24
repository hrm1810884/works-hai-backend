// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// GeneratePost implements POST /generate operation.
//
// Post id in storage to BE.
//
// POST /generate
func (UnimplementedHandler) GeneratePost(ctx context.Context, req *GeneratePostReq) (r GeneratePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// InitGet implements GET /init operation.
//
// Retrieve presigned URLs for both Human and AI drawings.
//
// GET /init
func (UnimplementedHandler) InitGet(ctx context.Context) (r InitGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RebaseDatabaseGet implements GET /rebase-database operation.
//
// Refresh position informations to clear data apparently.
//
// GET /rebase-database
func (UnimplementedHandler) RebaseDatabaseGet(ctx context.Context) (r RebaseDatabaseGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ViewGet implements GET /view operation.
//
// Viewer Page for human AI drawings.
//
// GET /view
func (UnimplementedHandler) ViewGet(ctx context.Context) (r ViewGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrRespStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrRespStatusCode) {
	r = new(ErrRespStatusCode)
	return r
}
