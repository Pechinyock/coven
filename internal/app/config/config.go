package config

import (
	"coven/internal/log"
	"time"
)

/* [TODO] Split config simaticly + strore everything inside this no matter what size it will be */
/* at startup will read configuration and check only required if we could run server and there's */
/* some configs that needs to be specified run initialize page only that will load this config */
/* and then helps to init application */

type CovenWebConfig struct {
	Address               string              `json:"address"`
	Port                  uint16              `json:"port"`
	ReadTimeout           time.Duration       `json:"readTimeout"`
	WriteTimeout          time.Duration       `json:"writeTimeout"`
	GloablLogConfig       log.GlobalLogConfig `json:"gloablLog"`
	EnableApiEndpoints    bool                `json:"enableApiEndpoints"`
	MiddlewaresConfigPath string              `json:"middlewaresConfigPath"`
	WebUIConfigPath       string              `json:"webUIConfigPath"`
	FileServerConfigPath  string              `json:"fileServerConfigPath"`
}

type MeddlewaresConfig struct {
	RequestLogger *RequestLoggerMiddlewareConfig `json:"requestLogger"`
}

type RequestLoggerMiddlewareConfig struct {
	Enabled  bool   `json:"enabled"`
	LogLevel string `json:"logLevel"`
}

type ShareDirConfig struct {
	RouteName   string `json:"routeName"`
	DirPath     string `json:"dirPath"`
	TokenSource string `jsong:"tokenSource"` // header, cookie, none
}

type FileServerConfig struct {
	ShareDirConfigs []ShareDirConfig `json:"shareDirConfigs"`
}
