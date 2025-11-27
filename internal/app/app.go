package app

import (
	"coven/internal/app/config"
	"fmt"
	"log/slog"
	"net/http"
	"time"
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
	withPermanent := addPermanentMiddlewares(router)
	withMiddlewares, err := addOptionalMiddlewares(withPermanent, conf.Middlewares)
	if err != nil {
		slog.Error("failed to register middlewares", "error message", err.Error())
		return err
	}

	err = loadUI(conf.WebUI)
	if err != nil {
		slog.Error("failed to load ui", "error message", err.Error())
		return err
	}

	serverAddress := fmt.Sprintf("%s:%d", conf.ServerOptions.Address, conf.ServerOptions.Port)
	readTimeout := time.Duration(conf.ServerOptions.ReadTimeoutSec) * time.Second
	writeTimeout := time.Duration(conf.ServerOptions.WriteTimeoutSec) * time.Second
	server := http.Server{
		Addr:         serverAddress,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      withMiddlewares,
	}
	if conf.Https != nil {
		slog.Info(fmt.Sprintf("running server at https://%s", serverAddress))
		err = server.ListenAndServeTLS(conf.Https.CertFilePath,
			conf.Https.CertKeyFilePath)
		if err != nil {
			slog.Error("failed to run http server", "error message", err.Error())
			return err
		}
	} else {
		slog.Warn("https certificate is not provided running as http")
		slog.Info(fmt.Sprintf("running server at http://%s", serverAddress))
		err = server.ListenAndServe()
		if err != nil {
			slog.Error("failed to run http server", "error message", err.Error())
			return err
		}
	}
	return nil
}
