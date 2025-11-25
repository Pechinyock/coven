package app

import (
	"fmt"
)

var (
	version     = "dev"
	gitShortSha = "none"
)

func Init() error {
	fmt.Printf("version: %s\ngit commit sha: %s\n", version, gitShortSha)
	config, err := readConfig()
	if err != nil {
		return err
	}
	err = config.Validate()
	if err != nil {
		return err
	}
	err = setupLog(config.Log)
	if err != nil {
		return err
	}
	return nil
}
