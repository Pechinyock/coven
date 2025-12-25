package cards

import (
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/utils"
	"encoding/base64"
	"encoding/json"
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

func GenerateCard(cardType, cardName, outputPath, templatesPath string, data any) error {
	if cardName == "" {
		return errors.New("card name could't be empty string")
	}
	jsonOutDir := shareddirs.CardsJsonDataDirPath.Path
	return saveCardData(jsonOutDir, cardType, cardName, data)
}

func saveCardData(cardDataDirPath, cardType, cardName string, data any) error {
	pathToTypeDir := path.Join(cardDataDirPath, cardType)
	if !utils.IsDirExists(pathToTypeDir) {
		if err := os.MkdirAll(pathToTypeDir, 0755); err != nil {
			return err
		}
	}
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s.json", cardName)
	saveToPath := filepath.Join(pathToTypeDir, filename)
	return os.WriteFile(saveToPath, jsonBytes, 0644)
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

func GetCardsFileNames(dirPath, fileType string) ([]string, error) {
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
