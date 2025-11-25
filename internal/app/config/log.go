package config

import "fmt"

type LogConfig struct {
	LogLevel string `json:"logLevel"`
	LogPath  string `json:"logPath"`
}

func (l *LogConfig) Validate() error {
	fmt.Println("log config validation is not implemented")
	return nil
}
