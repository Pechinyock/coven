package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type RequestLogger struct {
	logger slog.Logger
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func NewRequestLogger(logHandler slog.Handler) RequestLogger {
	lg := slog.New(logHandler)
	return RequestLogger{
		logger: *lg,
	}
}

func (l RequestLogger) AddFunc(router *http.ServeMux) (http.HandlerFunc, error) {
	result := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		router.ServeHTTP(rw, r)
		duration := time.Since(start)
		if rw.statusCode >= http.StatusOK && rw.statusCode < http.StatusMultipleChoices {
			l.logger.Info(r.RequestURI,
				"method", r.Method,
				"status code", rw.statusCode,
				"duration", duration,
			)
		} else if rw.statusCode >= http.StatusBadRequest && rw.statusCode <= http.StatusNetworkAuthenticationRequired {
			l.logger.Error(r.RequestURI,
				"method", r.Method,
				"status code", rw.statusCode,
				"duration", duration,
			)
		}

	})
	return result, nil
}
