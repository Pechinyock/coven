package config

import "fmt"

type RemoteStorageSettings struct {
	RepoStorageAddress string `json:"repoStorageAddress"`
	LocalDirPath       string `json:"localDirPath"`
	WorkingBranchName  string `json:"workingBranchName"`
	MainBranchName     string `json:"mainBranchName"`
}

func (c *RemoteStorageSettings) Validate() error {
	fmt.Println("remote storage config validation is not implemented")
	return nil
}
