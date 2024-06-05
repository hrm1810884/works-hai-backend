// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleApiKeyAuth handles ApiKeyAuth security.
	HandleApiKeyAuth(ctx context.Context, operationName string, t ApiKeyAuth) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityApiKeyAuth(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t ApiKeyAuth
	const parameterName = "X-Api-Key"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleApiKeyAuth(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// ApiKeyAuth provides ApiKeyAuth security value.
	ApiKeyAuth(ctx context.Context, operationName string) (ApiKeyAuth, error)
}

func (s *Client) securityApiKeyAuth(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.ApiKeyAuth(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"ApiKeyAuth\"")
	}
	req.Header.Set("X-Api-Key", t.APIKey)
	return nil
}
