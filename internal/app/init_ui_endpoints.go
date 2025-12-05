package app

import (
	"coven/internal/app/config"
	"coven/internal/endpoint/webui"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
)

func registerUIEndpoints(router *http.ServeMux, config *config.WebUIBundleConfig) error {
	if router == nil {
		return errors.New("can't add ui endpoints to nil router")
	}
	if config == nil {
		return errors.New("failed to add UI endpoints - configuration is empty")
	}
	uiEndpoints := webui.GetUIEndpoints()
	total := len(uiEndpoints)
	if total == 0 {
		slog.Warn("UI endpoint array is empty - no endpoints registered")
		return nil
	} else {
		slog.Debug(fmt.Sprintf("ready to register %d ui endpoints", total))
	}

	for _, e := range uiEndpoints {
		router.HandleFunc(e.Path, e.HandlerFunc)
		slog.Debug(fmt.Sprintf("ui endpoint has been added: %q", e.Path))
	}

	assetsPath := filepath.Join(config.RootPath, config.StaticFilesShare.DirPath)
	fs := http.FileServer(http.Dir(assetsPath))
	routeName := fmt.Sprintf("/%s/", config.StaticFilesShare.RouteName)
	router.Handle(routeName, http.StripPrefix(routeName, fs))

	physicalPath := filepath.Join(config.RootPath, config.StaticFilesShare.DirPath)
	slog.Debug("succesfully added ui static files routes",
		"physical dir path", physicalPath,
		"route path", config.StaticFilesShare.RouteName,
	)
	return nil
}
