package webui

import (
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/projection"
	"errors"
	"os"
	"path"
	"strings"
)

func loadImagesPrewiewData(poolGroupName string) ([]projection.ImageProj, error) {
	if poolGroupName == "" {
		return nil, errors.New("pool group name is empty")
	}
	pathToDir := path.Join(shareddirs.ImagePoolDirPath.Path, poolGroupName)
	files, err := os.ReadDir(pathToDir)
	if err != nil {
		return nil, err
	}
	var result []projection.ImageProj
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".png") {
			added := projection.ImageProj{
				FileName: file.Name(),
			}
			result = append(result, added)
		}
	}
	return result, nil
}
