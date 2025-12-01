package cards

import (
	"coven/internal/utils"
	"errors"
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
	return err
}
