package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hrm1810884/works-hai-backend/ogen"
)

func TestUploadURLGetHandler(t *testing.T) {
	uploadService := &MyUploadService{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		res, err := uploadService.UploadURLGet(ctx)
		if err != nil {
			http.Error(w, "Invalid input data", http.StatusBadRequest)
			return
		}

		switch res := res.(type) {
		case *ogen.UploadURLGetOK:
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(res); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
		case *ogen.UploadURLGetBadRequest:
			http.Error(w, "Invalid input data", http.StatusBadRequest)
		default:
			http.Error(w, "Unknown error", http.StatusInternalServerError)
		}
	}

	req, err := http.NewRequest("GET", "/upload-url", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler).ServeHTTP(rr, req)

	// ステータスコードのチェック
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// レスポンスボディのチェック
	expected := &ogen.UploadURLGetOK{
		PresignedURL: ogen.OptString{
			Value: "https://example-cloud-storage.com/user_drawing.png?signature=...",
			Set:   true,
		},
	}
	var actual ogen.UploadURLGetOK
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if actual.PresignedURL.Value != expected.PresignedURL.Value {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual.PresignedURL.Value, expected.PresignedURL.Value)
	}
}
