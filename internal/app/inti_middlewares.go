package app

import (
	"coven/internal/app/config"
	"coven/internal/app/middleware"
	"coven/internal/log"
	"net/http"
	"os"
)

func newRequestLogger(conf *config.RequestLoggerMiddlewareConfig) middleware.FuncMiddleware {
	lvl := parseLogLevel(conf.Level)
	logHandler := log.SimpleLogHandler{
		Level:  lvl,
		Writer: os.Stdout,
	}
	return middleware.NewRequestLogger(&logHandler)
}

func addOptionalMiddlewares(router *http.ServeMux, config *config.MeddlewaresConfig) error {
	if config.RequestLogger != nil {
		requestLogger := newRequestLogger(config.RequestLogger)
		requestLogger.AddFunc(router)
	}
	return nil
}
