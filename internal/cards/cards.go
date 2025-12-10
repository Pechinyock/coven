package cards

import (
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
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
	"characters":  "character_templ.html",
	"spells":      "",
	"secrets":     "",
	"curses":      "",
	"ingredients": "",
	"potions":     "",
}

func GenerateCard(cardType, cardName, outputPath, templatesPath string, data any) error {
	templateName := typeTemplPath[cardType]
	templatePath := filepath.Join(templatesPath, templateName)
	slog.Info("ready to generate card", "path", templateName)
	if templateName == "" {
		return errors.New("cant't find file")
	}
	rootTemplName := "card"
	templ, err := template.New(rootTemplName).ParseFiles(templatePath)
	if err != nil {
		return err
	}
	if !utils.IsDirExists(outputPath) {
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return err
		}
	}
	fullPath := path.Join(outputPath, cardType, cardName)
	if utils.IsFileExists(fullPath) {
		slog.Warn("overriding existg card", "path", fullPath)
	}
	fileResult, err := os.Create(fullPath + ".html")
	if err != nil {
		return err
	}
	defer fileResult.Close()
	err = templ.ExecuteTemplate(fileResult, templateName, data)
	if err != nil {
		return err
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
