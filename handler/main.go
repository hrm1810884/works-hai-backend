package handler

import (
	"context"
	"log"

	"github.com/hrm1810884/works-hai-backend/ogen"
)

type HaiHandler struct{}

func (*HaiHandler) ResourcePathPost(ctx context.Context, req *ogen.ResourcePathPostReq) (ogen.ResourcePathPostRes, error) {
	return &ogen.ResourcePathPostOK{}, nil
}

// NewError creates *ErrRespStatusCode from error returned by handler.
//
// Used for common default response.
func (*HaiHandler) NewError(ctx context.Context, err error) *ogen.ErrRespStatusCode {
	log.Fatalf("%v", err)
	return &ogen.ErrRespStatusCode{}
}
