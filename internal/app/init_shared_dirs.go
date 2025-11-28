package app

import (
	"coven/internal/app/config"
	"coven/internal/cards"
	"coven/internal/utils"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

func registerSharedDirs(router *http.ServeMux, conf *config.FileServerConfig) error {
	if router == nil {
		return errors.New("failed to register shared directories the router is nil")
	}
	if conf == nil {
		slog.Warn("file server configuration is not provided, setting it to defalt")
		conf = &config.FileServerConfig{
			CompleteCardsDir: defaultCompleteDir(),
			CardTemplatesDir: defaultTemplatesDir(),
			ImagePoolDir:     defaultImagePoolDir(),
		}
	}

	printWarn := func(dirName, path string) {
		path, err := utils.GetFullPath(path)
		if err != nil {
			return
		}
		label := fmt.Sprintf("%s directory is not provided setting it to default", dirName)
		slog.Warn(label, "\npath", path)
	}
	if conf.CompleteCardsDir == nil {
		conf.CompleteCardsDir = defaultCompleteDir()
		printWarn("complete cards", conf.CompleteCardsDir.DirPath)
	}
	if conf.CardTemplatesDir == nil {
		conf.CardTemplatesDir = defaultTemplatesDir()
		printWarn("card templates", conf.CardTemplatesDir.DirPath)
	}
	if conf.ImagePoolDir == nil {
		conf.ImagePoolDir = defaultImagePoolDir()
		printWarn("cards image pool", conf.ImagePoolDir.DirPath)
	}

	handlerSetter := func(routeName, path, source string) http.Handler {
		fs := http.FileServer(http.Dir(path))

		switch strings.ToLower(source) {
		case "header":
			panic("not implemented")
		case "cookie":
			panic("not implemented")
		case "none", "":
			{
				formated := fmt.Sprintf("/%s/", routeName)
				router.Handle(formated, http.StripPrefix(formated, fs))
				slog.Info("succesfly added directory as files server",
					"directory physical path", path,
					"route name", routeName,
				)
				return fs
			}
		default:
			panic(fmt.Sprintf("unknown token place for file server \n directory name: %s\n route name: %s", path, routeName))
		}
	}
	optional := conf.ShareDirConfigs
	if len(optional) != 0 {
		for _, e := range optional {
			handlerSetter(e.RouteName, e.DirPath, e.TokenSource)
		}
	} else {
		slog.Info("no additional share dir to registry")
	}
	handlerSetter(conf.CompleteCardsDir.RouteName, conf.CompleteCardsDir.DirPath, conf.CompleteCardsDir.TokenSource)
	handlerSetter(conf.CardTemplatesDir.RouteName, conf.CardTemplatesDir.DirPath, conf.CardTemplatesDir.TokenSource)
	handlerSetter(conf.ImagePoolDir.RouteName, conf.ImagePoolDir.DirPath, conf.ImagePoolDir.TokenSource)

	cards.CardsOutput = conf.CompleteCardsDir.DirPath
	cards.ImagePool = conf.ImagePoolDir.DirPath
	cards.CardTemplates = conf.CardTemplatesDir.DirPath

	return nil
}

func defaultCompleteDir() *config.ShareDirConfig {
	return &config.ShareDirConfig{
		RouteName:   "complete-cards",
		DirPath:     "./cards_complete",
		TokenSource: "none",
	}
}

func defaultTemplatesDir() *config.ShareDirConfig {
	return &config.ShareDirConfig{
		RouteName:   "card-templates",
		DirPath:     "./cards_templates",
		TokenSource: "none",
	}
}

func defaultImagePoolDir() *config.ShareDirConfig {
	return &config.ShareDirConfig{
		RouteName:   "image-pool",
		DirPath:     "./card_image_pool",
		TokenSource: "none",
	}
}
