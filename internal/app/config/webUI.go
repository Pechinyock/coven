package config

import "fmt"

type WebUIBundleConfig struct {
	RootPath         string          `json:"rootPath"`
	StaticFilesShare *ShareDirConfig `json:"staticFilesShare"`
	TemplatePaths    []string        `json:"templatePaths"`
}

func (w *WebUIBundleConfig) Validate() error {
	fmt.Println("web ui config validation is not implemented")
	return nil
}
