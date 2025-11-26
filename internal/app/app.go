package app

import (
	"coven/internal/app/config"
	"fmt"
	"net/http"
)

var (
	version     = "dev"
	gitShortSha = "none"
)

func Init() (*config.CovenWebConfig, error) {
	fmt.Printf("version: '%s' git commit sha: '%s'\n", version, gitShortSha)
	config, err := readConfig()
	if err != nil {
		return nil, err
	}
	err = config.Validate()
	if err != nil {
		return nil, err
	}
	setupLogger(config.Log)

	return config, nil
}

func Run(conf *config.CovenWebConfig) error {
	router := http.NewServeMux()

	err := addOptionalMiddlewares(router, conf.Middlewares)
	return err
}
