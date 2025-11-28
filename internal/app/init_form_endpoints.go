package app

import (
	"coven/internal/endpoint/form"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

func registerFormEndpoints(router *http.ServeMux) error {
	if router == nil {
		return errors.New("failed to register form endpoints router is nil")
	}
	endpoints := form.GetFormEndpoints()
	total := len(endpoints)
	if total == 0 {
		slog.Warn("no form endpoints provided")
		return nil
	} else {
		slog.Debug(fmt.Sprintf("ready to regestry %d form endpoints", total))
	}
	for _, e := range endpoints {
		router.HandleFunc(e.Path, e.HandlerFunc)
		slog.Debug(fmt.Sprintf("form endpoint has been added: %q", e.Path))
	}
	return nil
}
