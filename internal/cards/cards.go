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

var CardsOutput string
var ImagePool string
var CardTemplates string

func GenerateCard(cardType, cardName string, data any) error {
	templateName := typeTemplPath[cardType]
	templatePath := filepath.Join(CardTemplates, templateName)
	slog.Info("ready to generate card", "path", templateName)
	if templateName == "" {
		return errors.New("cant't find file")
	}
	rootTemplName := "card"
	templ, err := template.New(rootTemplName).ParseFiles(templatePath)
	if err != nil {
		return err
	}
	if !utils.IsDirExists(CardsOutput) {
		if err := os.MkdirAll(CardsOutput, 0755); err != nil {
			return err
		}
	}
	fullPath := path.Join(CardsOutput, cardType, cardName)
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
