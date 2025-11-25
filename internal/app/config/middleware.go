package config

import "fmt"

type RequestLoggerMiddlewareConfig struct {
	Enabled bool       `json:"enabled"`
	Log     *LogConfig `json:"options"`
}

type MeddlewaresConfig struct {
	RequestLogger *RequestLoggerMiddlewareConfig `json:"requestLogger"`
}

func (c *MeddlewaresConfig) Validate() error {
	fmt.Println("middleware config validation is not implemented")
	return nil
}
