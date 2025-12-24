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

func SaveCard(cardType, saveFormat, data string) error {
	if saveFormat == "png" {
		return savePng(cardType, data)
	}
	if saveFormat == "json" {
		return saveJson(cardType, data)
	}
	return nil
}

func saveJson(cardType, data string) error {
	fullPath := path.Join(shareddirs.CardsJsonDataDirPath.Path, cardType, "new_file.json")
	return os.WriteFile(fullPath, []byte(data), 0644)
}

func savePng(cardType, data string) error {
	pngData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	fullPath := path.Join(shareddirs.CompleteCardsDirPath.Path, cardType, "new_file.png")
	return os.WriteFile(fullPath, pngData, 0644)
}
