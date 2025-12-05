package app

import (
	"coven/internal/app/config"
	"coven/internal/log"
	"coven/internal/middleware"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func newRequestLogger(conf *config.RequestLoggerMiddlewareConfig) middleware.Middlware {
	if conf == nil {
		return nil
	}
	lvl := parseLogLevel(conf.Level)
	logHandler := log.ColoreLabelHandler{
		Level:  lvl,
		Writer: os.Stdout,
	}
	return middleware.NewRequestLogger(&logHandler)
}

/* [ISSUE] same func twice */
/* should have a 'global' map that difine permanent middlewares and it's order */
/* and add optional depend on config and then registy them all at once */
func addPermanentMiddlewares(router http.Handler) http.Handler {
	mdlw := map[string]middleware.Middlware{
		"recovery": middleware.ServerRecovery{},
	}
	for name, md := range mdlw {
		h, _ := md.Add(router)
		router = h
		slog.Info(fmt.Sprintf("middleware %q succesfuly added", name))
	}
	return router
}

func addOptionalMiddlewares(router http.Handler, config *config.MeddlewaresConfig) (http.Handler, error) {
	if config == nil {
		slog.Info("optional middlewares are disabled")
		return router, nil
	}
	mdlw := map[string]middleware.Middlware{
		"request logger": newRequestLogger(config.RequestLogger),
	}

	for name, md := range mdlw {
		if md == nil {
			slog.Info(fmt.Sprintf("%q is disabled", name))
			continue
		}
		h, err := md.Add(router)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to add %q middleware:\n   -> %s", name, err.Error()))
			continue
		}
		router = h
		slog.Info(fmt.Sprintf("middleware %q succesfuly added", name))
	}
	return router, nil
}
