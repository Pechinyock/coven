package cards

import (
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var CardTypes = map[string]string{
	"characters":  "Персонаж",
	"spells":      "Заклинание",
	"secrets":     "Секрет",
	"curses":      "Проклятье",
	"ingredients": "Ингредиент",
	"potions":     "Зелье",
}

func IsCardExists(cardType, cardName string) (bool, error) {
	pathToCondidate := filepath.Join(shareddirs.CardsJsonDataDirPath.Path,
		cardType,
		fmt.Sprintf("%s.json", cardName))

	files, err := filepath.Glob(pathToCondidate)
	if err != nil {
		return false, err
	}
	return len(files) > 0, nil
}

func GetCardFileNames(dirPath, fileType string) ([]string, error) {
	if fileType != "png" && fileType != "json" {
		return nil, errors.New("unknown file type")
	}
	files, err := filepath.Glob(filepath.Join(dirPath, fmt.Sprintf("*.%s", fileType)))
	if err != nil {
		return nil, err
	}
	var names []string
	for _, fullPath := range files {
		names = append(names, filepath.Base(fullPath))
	}
	return names, nil
}

func SaveCard(cardType, cardName, saveFormat, data string) error {
	if saveFormat == "png" {
		return savePng(cardType, cardName, data)
	}
	if saveFormat == "json" {
		return saveJson(cardType, cardName, data)
	}
	return nil
}

func DeleteCard(cardType, cardName string) error {
	pathToCondidate := filepath.Join(shareddirs.CardsJsonDataDirPath.Path,
		cardType,
		fmt.Sprintf("%s.json", cardName),
	)

	files, err := filepath.Glob(pathToCondidate)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("coudn't find card json data, type %q, name %q", cardType, cardName)
	}
	if len(files) > 1 {
		return errors.New("file glob returns more than 1 value")
	}
	err = os.Remove(files[0])
	if err != nil {
		return err
	}
	pathToCondidate = filepath.Join(shareddirs.CompleteCardsDirPath.Path,
		cardType,
		fmt.Sprintf("%s.png", cardName),
	)
	files, err = filepath.Glob(pathToCondidate)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("coudn't find card png image, type %q, name %q", cardType, cardName)
	}
	if len(files) > 1 {
		return errors.New("file glob returns more than 1 value")
	}
	err = os.Remove(files[0])
	if err != nil {
		return err
	}
	return nil
}

func LoadCardJsonData(cardType, cardName string) ([]byte, error) {
	pathToJson := path.Join(shareddirs.CardsJsonDataDirPath.Path,
		cardType,
		fmt.Sprintf("%s.json", cardName),
	)
	if !utils.IsFileExists(pathToJson) {
		return nil, fmt.Errorf("file not found: %s", pathToJson)
	}
	data, err := os.ReadFile(pathToJson)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func saveJson(cardType, cardName, data string) error {
	fullPath := path.Join(shareddirs.CardsJsonDataDirPath.Path, cardType, fmt.Sprintf("%s.json", cardName))
	return os.WriteFile(fullPath, []byte(data), 0644)
}

func savePng(cardType, cardName, data string) error {
	pngData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	fullPath := path.Join(shareddirs.CompleteCardsDirPath.Path, cardType, fmt.Sprintf("%s.png", cardName))
	return os.WriteFile(fullPath, pngData, 0644)
}
