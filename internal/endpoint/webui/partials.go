package webui

import (
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/projection"
	"os"
	"path"
	"strings"
)

func loadPartialsPreview() ([]projection.ImageProj, error) {
	pathToPartials := path.Join(shareddirs.PartialsDirPath.Path, "preview")
	files, err := os.ReadDir(pathToPartials)
	if err != nil {
		return nil, err
	}
	var result []projection.ImageProj
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".png") {
			partialName := strings.TrimSuffix(file.Name(), ".png")
			added := projection.ImageProj{
				FileName: partialName,
			}
			result = append(result, added)
		}
	}
	return result, nil
}
