package cards

import (
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/utils"
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

var typeTemplPath = map[string]string{
	"characters":  "character_card",
	"spells":      "",
	"secrets":     "",
	"curses":      "",
	"ingredients": "",
	"potions":     "",
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
