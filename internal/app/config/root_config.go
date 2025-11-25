package config

import (
	"errors"
	"fmt"
)

type CovenWebConfig struct {
	ServerOptions *ServerConfig      `json:"server"`
	Https         *HttpsConfig       `json:"https"`
	Log           *LogConfig         `json:"log"`
	Middlewares   *MeddlewaresConfig `json:"middlewares"`
	WebUI         *WebUIBundleConfig `json:"webUI"`
	FileServer    *FileServerConfig  `json:"fileServer"`
}

func (c *CovenWebConfig) Validate() error {
	fmt.Println("root config validation is not implemented")
	childs := []Config{
		c.ServerOptions,
		c.Https,
		c.Log,
		c.Middlewares,
		c.WebUI,
		c.FileServer,
	}
	errs := []error{}
	for _, cfg := range childs {
		// if cfg == nil {
		// 	continue
		// }
		err := cfg.Validate()
		if err != nil {
			errs = append(errs, err)
		}
	}

	multiErr := errors.Join(errs...)
	return multiErr
}
