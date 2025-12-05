package middleware

import (
	"log/slog"
	"net/http"
)

type AuthMiddleware struct {
	authFailed http.HandlerFunc
}

func NewHeaderTokenAuth() *AuthMiddleware {
	return &AuthMiddleware{
		authFailed: func(w http.ResponseWriter, r *http.Request) {
			slog.Error("auth failed")
		},
	}
}

func NewCookieTokenAuth() *AuthMiddleware {
	return &AuthMiddleware{
		authFailed: func(w http.ResponseWriter, r *http.Request) {
			slog.Error("auth failed")
		},
	}
}

func (a AuthMiddleware) Add(next http.Handler) (http.Handler, error) {
	result := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("auth triggered")
		next.ServeHTTP(w, r)
	})
	return result, nil
}
