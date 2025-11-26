package config

import "fmt"

type RequestLoggerMiddlewareConfig struct {
	Level string `json:"level"`
}

type MeddlewaresConfig struct {
	RequestLogger *RequestLoggerMiddlewareConfig `json:"requestLogger"`
}

func (c *MeddlewaresConfig) Validate() error {
	fmt.Println("middleware config validation is not implemented")
	return nil
}
