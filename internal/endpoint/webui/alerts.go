package webui

import (
	"coven/internal/projection"
	"log/slog"
	"net/http"
)

const (
	alertTemplName = "alert"
	successType    = "success"
	failedType     = "danger"
)

func SendSucces(w http.ResponseWriter, message string) {
	if message == "" {
		slog.Warn("trying to send an empty success alert")
	}
	UIBundle.Render(alertTemplName, w, projection.AlertProj{
		Type:    successType,
		Message: message,
	})
}

func SendFailed(w http.ResponseWriter, message string) {
	if message == "" {
		slog.Warn("trying to send an empty failed alert")
	}
	UIBundle.Render(alertTemplName, w, projection.AlertProj{
		Type:    failedType,
		Message: message,
	})
}
