package app

import (
	"coven/internal/app/config"
	"coven/internal/endpoint"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

var (
	Version     = "dev"
	GitShortSha = "none"
)

func Init() (*config.CovenWebConfig, error) {
	fmt.Printf("version: '%s' git commit sha: '%s'\n", Version, GitShortSha)
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

	err = registerUIEndpoints(router, conf.WebUI)
	if err != nil {
		slog.Error("failed to register ui endpoints", "error message", err.Error())
		return err
	}

	err = configureRemoteRepo(conf.RemoteStorage)
	if err != nil {
		slog.Error("failed to configure remote repo", "error message", err.Error())
		return err
	}

	err = registerSharedDirs(router, conf.FileServer)
	if err != nil {
		slog.Error("failed to register shared directories", "error message", err.Error())
		return err
	}

	err = registerFormEndpoints(router)
	if err != nil {
		slog.Error("failed to register form endpoints", "error message", err.Error())
		return err
	}

	serverAddress := fmt.Sprintf("%s:%d", conf.ServerOptions.Address, conf.ServerOptions.Port)

	endpoint.Address = conf.ServerOptions.Address
	endpoint.Port = conf.ServerOptions.Port

	readTimeout := time.Duration(conf.ServerOptions.ReadTimeoutSec) * time.Second
	writeTimeout := time.Duration(conf.ServerOptions.WriteTimeoutSec) * time.Second
	server := http.Server{
		Addr:         serverAddress,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      withMiddlewares,
	}
	if conf.Https != nil {
		scheme := "https"
		endpoint.Scheme = scheme

		address := fmt.Sprintf("%s://%s", scheme, serverAddress)
		slog.Info(fmt.Sprintf("running server at %s", address))
		err = server.ListenAndServeTLS(conf.Https.CertFilePath,
			conf.Https.CertKeyFilePath)
		if err != nil {
			slog.Error("failed to run http server", "error message", err.Error())
			return err
		}
	} else {
		slog.Warn("https certificate is not provided running as http")
		scheme := "http"
		endpoint.Scheme = scheme

		address := fmt.Sprintf("%s://%s", scheme, serverAddress)
		slog.Info(fmt.Sprintf("running server at %s", address))
		err = server.ListenAndServe()
		if err != nil {
			slog.Error("failed to run http server", "error message", err.Error())
			return err
		}
	}
	return nil
}
