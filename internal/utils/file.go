package utils

import (
	"os"
	"path/filepath"
	"strings"
)

const pathGlobSymbols = "*?["

func GetFullPath(path string) (string, error) {
	var fullFilePath string
	if filepath.IsAbs(path) {
		fullFilePath = path
	} else {
		full, err := filepath.Abs(path)
		if err != nil {
			return "", err
		}
		fullFilePath = full
	}
	return fullFilePath, nil
}

func IsFileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || fileInfo.IsDir() {
		return false
	}
	return true
}

func IsDirExists(path string) bool {
	dirInfo, err := os.Stat(path)
	if err != nil || !dirInfo.IsDir() {
		return false
	}
	return true
}

func IsGlob(path string) bool {
	return strings.ContainsAny(path, pathGlobSymbols)
}

func GetFileName(path string, withExt bool) string {
	fileName := filepath.Base(path)
	if withExt {
		return fileName
	} else {
		return strings.TrimSuffix(fileName, filepath.Ext(fileName))
	}
}

func IsFilePath(path string) bool {
	base := filepath.Base(path)
	ext := filepath.Ext(base)

	return ext != "" && base != ext
}

func IsExtension(fileName, ext string) bool {
	return strings.EqualFold(fileName, ext)
}

func IsValidPath(s string) bool {
	return s == filepath.Clean(s) && filepath.Base(s) == s
}
