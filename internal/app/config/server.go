package config

import "fmt"

type ServerConfig struct {
	Address         string `json:"address"`
	Port            uint16 `json:"port"`
	ReadTimeoutSec  uint16 `json:"readTimeoutSec"`
	WriteTimeoutSec uint16 `json:"writeTimeoutSec"`
}

func (s *ServerConfig) Validate() error {
	fmt.Println("server config validation is not implemented")
	return nil
}
