package form

import (
	"coven/internal/cards"
	"coven/internal/endpoint/webui"
	"coven/internal/utils"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

func cardHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			cardType := r.URL.Query().Get("cardType")
			if cardType == "" {
				slog.Error("failed to load json card data, card type is empty")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			cardName := r.URL.Query().Get("cardName")
			if cardName == "" {
				slog.Error("failed to load json card data, card name is empty")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			data, err := cards.LoadCardJsonData(cardType, cardName)
			if err != nil {
				slog.Error("failed to load card json data",
					"card name", cardName,
					"card type", cardType,
					"error message", err.Error(),
				)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if len(data) == 0 {
				slog.Error("failed to load card json data, data is empty")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	case http.MethodPost:
		{
			cardType := r.FormValue("cardType")
			if len(cardType) == 0 {
				webui.SendFailed(w, "Нужно выбрать тип карты")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			jsonCardData := r.FormValue("jsonData")
			if jsonCardData == "" {
				webui.SendFailed(w, "json не может быть пустым")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			pngCardData := r.FormValue("pngData")
			if pngCardData == "" {
				webui.SendFailed(w, "png не может быть пустым")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			name := r.FormValue("cardName")
			if !utils.IsInValidLenghtRange(name) {
				webui.SendFailed(w, "Название не может быть пустым или содержать больше 247 символов")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !utils.IsValidFileName(name) {
				webui.SendFailed(w, `Название карты содержит недопустимые символы: \/:*?"<>|`)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			completeName := utils.ReplaceAllSpaces(name)
			isExists, err := cards.IsCardExists(cardType, completeName)
			if err != nil {
				webui.SendFailed(w, "При сохранении произошла внутренняя ошибка сервера")
				slog.Error("failed to check if card exists", "card name", name)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if isExists {
				webui.SendFailed(w, fmt.Sprintf("карта с названием %q уже существует", name))
				webui.UIBundle.Render("override_card", w, nil)
				w.WriteHeader(http.StatusConflict)
				return
			}
			err = cards.SaveCard(cardType, completeName, "json", jsonCardData)
			if err != nil {
				webui.SendFailed(w, fmt.Sprintf("failed to save json card data %s", err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = cards.SaveCard(cardType, completeName, "png", pngCardData)
			if err != nil {
				webui.SendFailed(w, fmt.Sprintf("failed to save png card data %s", err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			webui.SendSucces(w, fmt.Sprintf("Карта успешно %q создана", name))
		}
	case http.MethodPatch:
		{
			cardType := r.FormValue("cardType")
			if len(cardType) == 0 {
				webui.SendFailed(w, "Нужно выбрать тип карты")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			jsonCardData := r.FormValue("jsonData")
			if jsonCardData == "" {
				webui.SendFailed(w, "json не может быть пустым")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			pngCardData := r.FormValue("pngData")
			if pngCardData == "" {
				webui.SendFailed(w, "png не может быть пустым")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			name := r.FormValue("cardName")
			if !utils.IsInValidLenghtRange(name) {
				webui.SendFailed(w, "Название не может быть пустым или содержать больше 247 символов")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !utils.IsValidFileName(name) {
				webui.SendFailed(w, `Название карты содержит недопустимые символы: \/:*?"<>|`)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			completeName := utils.ReplaceAllSpaces(name)
			isExists, err := cards.IsCardExists(cardType, completeName)
			if err != nil {
				webui.SendFailed(w, "При сохранении произошла внутренняя ошибка сервера")
				slog.Error("failed to check if card exists", "card name", name)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if !isExists {
				webui.SendFailed(w, "Пока ты думал, блять, карту уже удалили... жми просто сохранить")
				w.WriteHeader(http.StatusConflict)
				return
			}
			err = cards.SaveCard(cardType, completeName, "json", jsonCardData)
			if err != nil {
				webui.SendFailed(w, fmt.Sprintf("failed to save card %s", err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = cards.SaveCard(cardType, completeName, "png", pngCardData)
			if err != nil {
				webui.SendFailed(w, fmt.Sprintf("failed to save card %s", err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			webui.SendSucces(w, fmt.Sprintf("Карта успешно %q перезаписана", name))
		}
	case http.MethodDelete:
		{
			body, err := io.ReadAll(r.Body)
			if err != nil {
				webui.SendFailed(w, "Произошла внутренняя ошибка сервера")
				slog.Error("failed to parse body request", "error message", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			values, err := url.ParseQuery(string(body))
			if err != nil {
				webui.SendFailed(w, "Произошла внутренняя ошибка сервера")
				slog.Error("failed to parse query", "error message", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			cardType := values.Get("deleteCardType")
			if cardType == "" {
				webui.SendFailed(w, "Не удалось удалить карту тип не может быть пустым")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			_, exists := cards.CardTypes[cardType]
			if !exists {
				webui.SendFailed(w, fmt.Sprintf("Не удалось удалить карту неизвестный тип карты %q", cardType))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			cardName := values.Get("deleteCardName")
			if cardName == "" {
				webui.SendFailed(w, "Не удалось удалить карту название не может быть пустым")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = cards.DeleteCard(cardType, cardName)
			if err != nil {
				webui.SendFailed(w, "Произошла внутренняя ошибка сервера")
				slog.Error("card deletion failed", "error message", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			webui.SendSucces(w, fmt.Sprintf("Карта %q была удалена", cardName))
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
