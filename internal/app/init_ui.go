package app

import (
	"coven/internal/app/config"
	"coven/internal/endpoint/webui"
	"coven/internal/ui"
	"errors"
)

func loadUI(conf *config.WebUIBundleConfig) error {
	if conf == nil {
		return errors.New("UI config is not specified")
	}
	bdl, err := ui.New(conf)
	if err != nil {
		return nil
	}
	webui.SetUIBundle(*bdl)
	return nil
}
