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

type GeneratePostBadRequest struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *GeneratePostBadRequest) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *GeneratePostBadRequest) SetError(val OptString) {
	s.Error = val
}

func (*GeneratePostBadRequest) generatePostRes() {}

type GeneratePostOK struct {
	URL OptString `json:"url"`
}

// GetURL returns the value of URL.
func (s *GeneratePostOK) GetURL() OptString {
	return s.URL
}

// SetURL sets the value of URL.
func (s *GeneratePostOK) SetURL(val OptString) {
	s.URL = val
}

func (*GeneratePostOK) generatePostRes() {}

type GeneratePostReq struct {
	// User id of the experience.
	UserId string `json:"userId"`
}

// GetUserId returns the value of UserId.
func (s *GeneratePostReq) GetUserId() string {
	return s.UserId
}

// SetUserId sets the value of UserId.
func (s *GeneratePostReq) SetUserId(val string) {
	s.UserId = val
}

// InitGetBadRequest is response for InitGet operation.
type InitGetBadRequest struct{}

func (*InitGetBadRequest) initGetRes() {}

// InitGetInternalServerError is response for InitGet operation.
type InitGetInternalServerError struct{}

func (*InitGetInternalServerError) initGetRes() {}

type InitGetOK struct {
	Result InitGetOKResult `json:"result"`
}

// GetResult returns the value of Result.
func (s *InitGetOK) GetResult() InitGetOKResult {
	return s.Result
}

// SetResult sets the value of Result.
func (s *InitGetOK) SetResult(val InitGetOKResult) {
	s.Result = val
}

func (*InitGetOK) initGetRes() {}

type InitGetOKResult struct {
	// Presigned URL for human drawing upload.
	HumanDrawing  string    `json:"humanDrawing"`
	TopDrawing    OptString `json:"topDrawing"`
	RightDrawing  OptString `json:"rightDrawing"`
	BottomDrawing OptString `json:"bottomDrawing"`
	LeftDrawing   OptString `json:"leftDrawing"`
}

// GetHumanDrawing returns the value of HumanDrawing.
func (s *InitGetOKResult) GetHumanDrawing() string {
	return s.HumanDrawing
}

// GetTopDrawing returns the value of TopDrawing.
func (s *InitGetOKResult) GetTopDrawing() OptString {
	return s.TopDrawing
}

// GetRightDrawing returns the value of RightDrawing.
func (s *InitGetOKResult) GetRightDrawing() OptString {
	return s.RightDrawing
}

// GetBottomDrawing returns the value of BottomDrawing.
func (s *InitGetOKResult) GetBottomDrawing() OptString {
	return s.BottomDrawing
}

// GetLeftDrawing returns the value of LeftDrawing.
func (s *InitGetOKResult) GetLeftDrawing() OptString {
	return s.LeftDrawing
}

// SetHumanDrawing sets the value of HumanDrawing.
func (s *InitGetOKResult) SetHumanDrawing(val string) {
	s.HumanDrawing = val
}

// SetTopDrawing sets the value of TopDrawing.
func (s *InitGetOKResult) SetTopDrawing(val OptString) {
	s.TopDrawing = val
}

// SetRightDrawing sets the value of RightDrawing.
func (s *InitGetOKResult) SetRightDrawing(val OptString) {
	s.RightDrawing = val
}

// SetBottomDrawing sets the value of BottomDrawing.
func (s *InitGetOKResult) SetBottomDrawing(val OptString) {
	s.BottomDrawing = val
}

// SetLeftDrawing sets the value of LeftDrawing.
func (s *InitGetOKResult) SetLeftDrawing(val OptString) {
	s.LeftDrawing = val
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
