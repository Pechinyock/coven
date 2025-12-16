package app

import (
	"coven/internal/app/config"
	"coven/internal/cards"
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/utils"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
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
			CardsJsonDataDir: defaultJsonDataDir(),
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
	if conf.CardsJsonDataDir == nil {
		conf.CardsJsonDataDir = defaultJsonDataDir()
		printWarn("cards json data", conf.CardsJsonDataDir.DirPath)
	}
	covenDirs := map[string]*config.ShareDirConfig{
		"image_pool":      conf.ImagePoolDir,
		"cards_json_data": conf.CardsJsonDataDir,
		"cards_templates": conf.CardTemplatesDir,
		"complete_cards":  conf.CompleteCardsDir,
		"card_styles":     conf.CardStylesDir,
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
			err := utils.CreateDirIfNotExists(e.DirPath)
			if err != nil {
				slog.Error("failed to create optional shared dir", "path", e.DirPath)
			}
		}
	} else {
		slog.Info("no additional share dir to registry")
	}

	for _, dir := range covenDirs {
		handlerSetter(dir.RouteName, dir.DirPath, dir.TokenSource)
	}

	complete := shareddirs.SharedDirPaths{
		Path: conf.CompleteCardsDir.DirPath,
		Uri:  conf.CompleteCardsDir.RouteName,
	}
	shareddirs.CompleteCardsDirPath = complete

	jsonData := shareddirs.SharedDirPaths{
		Path: conf.CardsJsonDataDir.DirPath,
		Uri:  conf.CardsJsonDataDir.RouteName,
	}
	shareddirs.CardsJsonDataDirPath = jsonData

	imgPool := shareddirs.SharedDirPaths{
		Path: conf.ImagePoolDir.DirPath,
		Uri:  conf.ImagePoolDir.RouteName,
	}
	shareddirs.ImagePoolDirPath = imgPool

	templates := shareddirs.SharedDirPaths{
		Path: conf.CardTemplatesDir.DirPath,
		Uri:  conf.CardTemplatesDir.RouteName,
	}
	shareddirs.CardTemplatesDirPath = templates

	cardStyles := shareddirs.SharedDirPaths{
		Path: conf.CardStylesDir.DirPath,
		Uri:  conf.CardStylesDir.RouteName,
	}
	shareddirs.CardStylesDirPath = cardStyles

	createSubDirs := func(base string, names []string) error {
		for _, e := range names {
			fullPath := filepath.Join(base, e)
			err := utils.CreateDirIfNotExists(fullPath)
			if err != nil {
				return err
			}
		}
		return nil
	}

	names := []string{}
	for n := range cards.CardTypes {
		names = append(names, n)
	}

	/* [ISSUE] Could be a function itterating through map defined at top of this func */
	err := createSubDirs(complete.Path, names)
	if err != nil {
		return err
	}

	err = createSubDirs(jsonData.Path, names)
	if err != nil {
		return err
	}

	err = createSubDirs(imgPool.Path, names)
	if err != nil {
		return err
	}

	err = createSubDirs(cardStyles.Path, names)
	if err != nil {
		return err
	}

	logSetup := func(name string, elem shareddirs.SharedDirPaths) {
		slog.Info(fmt.Sprintf("%s directory has been initialized", name),
			"physical path", elem.Path,
			"web uri", elem.Uri,
		)
	}

	err = utils.CreateDirIfNotExists(complete.Path)
	if err != nil {
		return err
	}

	err = utils.CreateDirIfNotExists(jsonData.Path)
	if err != nil {
		return err
	}

	err = utils.CreateDirIfNotExists(templates.Path)
	if err != nil {
		return err
	}

	err = utils.CreateDirIfNotExists(imgPool.Path)
	if err != nil {
		return err
	}

	err = utils.CreateDirIfNotExists(cardStyles.Path)
	if err != nil {
		return err
	}

	logSetup("complete cards", complete)
	logSetup("cards json data", jsonData)
	logSetup("card templates", templates)
	logSetup("image pool", imgPool)
	logSetup("card styles", cardStyles)

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

func defaultJsonDataDir() *config.ShareDirConfig {
	return &config.ShareDirConfig{
		RouteName:   "cards-json-data",
		DirPath:     "./cards_json_data",
		TokenSource: "none",
	}
}
