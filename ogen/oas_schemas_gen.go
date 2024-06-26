// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"
)

func (s *ErrRespStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type ApiKeyAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *ApiKeyAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *ApiKeyAuth) SetAPIKey(val string) {
	s.APIKey = val
}

type ErrResp struct {
	// A detailed error message.
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *ErrResp) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *ErrResp) SetError(val OptString) {
	s.Error = val
}

// ErrRespStatusCode wraps ErrResp with StatusCode.
type ErrRespStatusCode struct {
	StatusCode int
	Response   ErrResp
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrRespStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrRespStatusCode) GetResponse() ErrResp {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrRespStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrRespStatusCode) SetResponse(val ErrResp) {
	s.Response = val
}

type ImageGenerationPostBadRequest struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *ImageGenerationPostBadRequest) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *ImageGenerationPostBadRequest) SetError(val OptString) {
	s.Error = val
}

func (*ImageGenerationPostBadRequest) imageGenerationPostRes() {}

type ImageGenerationPostOK struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *ImageGenerationPostOK) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *ImageGenerationPostOK) SetMessage(val OptString) {
	s.Message = val
}

func (*ImageGenerationPostOK) imageGenerationPostRes() {}

type ImageGenerationPostReq struct {
	// The x-coordinate.
	X OptInt `json:"x"`
	// The y-coordinate.
	Y OptInt `json:"y"`
}

// GetX returns the value of X.
func (s *ImageGenerationPostReq) GetX() OptInt {
	return s.X
}

// GetY returns the value of Y.
func (s *ImageGenerationPostReq) GetY() OptInt {
	return s.Y
}

// SetX sets the value of X.
func (s *ImageGenerationPostReq) SetX(val OptInt) {
	s.X = val
}

// SetY sets the value of Y.
func (s *ImageGenerationPostReq) SetY(val OptInt) {
	s.Y = val
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// PresignedUrlsGetBadRequest is response for PresignedUrlsGet operation.
type PresignedUrlsGetBadRequest struct{}

func (*PresignedUrlsGetBadRequest) presignedUrlsGetRes() {}

// PresignedUrlsGetInternalServerError is response for PresignedUrlsGet operation.
type PresignedUrlsGetInternalServerError struct{}

func (*PresignedUrlsGetInternalServerError) presignedUrlsGetRes() {}

type PresignedUrlsGetOK struct {
	Result PresignedUrlsGetOKResult `json:"result"`
}

// GetResult returns the value of Result.
func (s *PresignedUrlsGetOK) GetResult() PresignedUrlsGetOKResult {
	return s.Result
}

// SetResult sets the value of Result.
func (s *PresignedUrlsGetOK) SetResult(val PresignedUrlsGetOKResult) {
	s.Result = val
}

func (*PresignedUrlsGetOK) presignedUrlsGetRes() {}

type PresignedUrlsGetOKResult struct {
	// Presigned URL for human drawing upload.
	HumanDrawing  string    `json:"humanDrawing"`
	TopDrawing    OptString `json:"topDrawing"`
	RightDrawing  OptString `json:"rightDrawing"`
	BottomDrawing OptString `json:"bottomDrawing"`
	LeftDrawing   OptString `json:"leftDrawing"`
}

// GetHumanDrawing returns the value of HumanDrawing.
func (s *PresignedUrlsGetOKResult) GetHumanDrawing() string {
	return s.HumanDrawing
}

// GetTopDrawing returns the value of TopDrawing.
func (s *PresignedUrlsGetOKResult) GetTopDrawing() OptString {
	return s.TopDrawing
}

// GetRightDrawing returns the value of RightDrawing.
func (s *PresignedUrlsGetOKResult) GetRightDrawing() OptString {
	return s.RightDrawing
}

// GetBottomDrawing returns the value of BottomDrawing.
func (s *PresignedUrlsGetOKResult) GetBottomDrawing() OptString {
	return s.BottomDrawing
}

// GetLeftDrawing returns the value of LeftDrawing.
func (s *PresignedUrlsGetOKResult) GetLeftDrawing() OptString {
	return s.LeftDrawing
}

// SetHumanDrawing sets the value of HumanDrawing.
func (s *PresignedUrlsGetOKResult) SetHumanDrawing(val string) {
	s.HumanDrawing = val
}

// SetTopDrawing sets the value of TopDrawing.
func (s *PresignedUrlsGetOKResult) SetTopDrawing(val OptString) {
	s.TopDrawing = val
}

// SetRightDrawing sets the value of RightDrawing.
func (s *PresignedUrlsGetOKResult) SetRightDrawing(val OptString) {
	s.RightDrawing = val
}

// SetBottomDrawing sets the value of BottomDrawing.
func (s *PresignedUrlsGetOKResult) SetBottomDrawing(val OptString) {
	s.BottomDrawing = val
}

// SetLeftDrawing sets the value of LeftDrawing.
func (s *PresignedUrlsGetOKResult) SetLeftDrawing(val OptString) {
	s.LeftDrawing = val
}
