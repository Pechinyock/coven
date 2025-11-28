package config

import "fmt"

type ShareDirConfig struct {
	RouteName   string `json:"routeName"`
	DirPath     string `json:"dirPath"`
	TokenSource string `json:"tokenSource"` // header, cookie, none
}

type FileServerConfig struct {
	ShareDirConfigs  []ShareDirConfig `json:"sharedDirs"`
	CompleteCardsDir *ShareDirConfig  `json:"completeCardsDir"`
	CardTemplatesDir *ShareDirConfig  `json:"cardTemplatesDir"`
	ImagePoolDir     *ShareDirConfig  `json:"imagePoolDir"`
}

func (f *FileServerConfig) Validate() error {
	fmt.Println("file server config validation is not implemented")
	return nil
}
