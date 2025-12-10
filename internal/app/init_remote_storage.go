package app

import (
	"coven/internal/app/config"
	"coven/internal/git"
	"coven/internal/utils"
	"fmt"
	"log/slog"
	"os/exec"
	"path"
)

func configureRemoteRepo(conf *config.RemoteStorageSettings) error {
	if conf == nil {
		slog.Info("no settings provided for remote storage")
		return nil
	}
	physPath, err := utils.GetFullPath(conf.LocalDirPath)
	if err != nil {
		return err
	}
	if !utils.IsDirExists(physPath) {
		cmd := exec.Command("git", "clone", conf.RepoStorageAddress, conf.LocalDirPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		slog.Info(fmt.Sprintf("succesfuly cloned: %s", output))
	} else if !utils.IsDirExists(path.Join(physPath, ".git")) {
		slog.Error("rempote storage directory already exists, but '.git' folder doesn't")
		slog.Error(fmt.Sprintf("choose another direcotry or delete the existing one %s", physPath))
	} else {
		slog.Info("remote storage already exists")
	}

	git.GitDirPath = conf.LocalDirPath
	git.GitOriginUrl = conf.RepoStorageAddress

	return nil
}
